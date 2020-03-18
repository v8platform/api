package pool

import (
	"context"
	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"
)

type transferFn func(ctx context.Context, client *sftp.Client, src, dest string) error

type WorkerPool struct {
	maxSize int
	ctx     context.Context
	cancel  func()
	conn    *ssh.Client

	// Protects access to fields below
	mu      sync.Mutex
	wg      sync.WaitGroup
	pool    chan *sftp.Client
	running chan *sftp.Client

	download transferFn
	upload   transferFn

	cls []closer
	err *multierror.Error
}

type Option func(wp *WorkerPool)
type closer func() error

func WithMaxSize(max int) Option {

	return func(wp *WorkerPool) {
		wp.maxSize = max
	}

}

func WithContext(ctx context.Context) Option {

	return func(wp *WorkerPool) {
		wp.ctx = ctx
	}

}

func WithDownload(fn transferFn) Option {

	return func(wp *WorkerPool) {
		wp.download = fn
	}

}

func WithUpload(fn transferFn) Option {

	return func(wp *WorkerPool) {
		wp.upload = fn
	}

}

type directionType uint

const (
	downloadType directionType = iota
	uploadType
)

func (t directionType) String() string {

	switch t {

	case downloadType:
		return "download"
	case uploadType:
		return "upload"
	default:
		return "unknown"
	}

}

type fileTransfer struct {
	direction directionType
	src       string
	dest      string
}

func downloadFile(ctx context.Context, client *sftp.Client, src, dest string) error {

	targetDir := filepath.Dir(dest)
	targetDir = filepath.ToSlash(targetDir)

	//size := int64(0)

	err := os.MkdirAll(targetDir, os.ModePerm)

	if err != nil {
		log.Fatal(err)
	}

	// create destination file
	dstFile, err := os.Create(dest)
	if err != nil {
		log.Fatal(err)
	}
	defer dstFile.Close()

	// create source file
	srcFile, err := client.Open(src)
	if err != nil {
		log.Fatal(err)
	}
	defer srcFile.Close()

	// copy source file to destination file
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		log.Fatal(err)
	}

	return err

}

func transferWithContext(ctx context.Context, fn transferFn, client *sftp.Client, src, dest string) error {

	// Создаем канал
	done := make(chan bool, 1)

	var err error
	// Запускаем выполнение медленной задачи в горутине
	// Передаем канал для коммуникаций
	go func() {
		err = fn(ctx, client, src, dest)
		done <- true
	}()

	// Используем select для выхода по истечении времени жизни контекста
	select {
	case <-ctx.Done():
		// Если контекст отменен, выбирается этот случай
		// Это случается, если вызвали cancel
		return ctx.Err()

	case <-done:
		// Этот вариант выбирается, когда работа завершается до отмены контекста
	}

	return err

}

func uploadFile(ctx context.Context, client *sftp.Client, src, dest string) error {

	targetDir := filepath.Dir(dest)
	targetDir = filepath.ToSlash(targetDir)

	//size := int64(0)

	err := os.MkdirAll(targetDir, os.ModePerm)

	if err != nil {
		log.Fatal(err)
	}

	// create destination file
	dstFile, err := os.Create(dest)
	if err != nil {
		log.Fatal(err)
	}
	defer dstFile.Close()

	// create source file
	srcFile, err := client.Open(src)
	if err != nil {
		log.Fatal(err)
	}
	defer srcFile.Close()

	// copy source file to destination file
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		log.Fatal(err)
	}

	return err

}

// NewPool creates a new pool of connections and starts GC. If no configuration
// is specified (nil), defaults values are used.
func NewPool(conn *ssh.Client, opts ...Option) *WorkerPool {

	p := &WorkerPool{
		conn:     conn,
		maxSize:  1,
		ctx:      context.Background(),
		download: downloadFile,
		upload:   uploadFile,
	}

	for _, opt := range opts {
		opt(p)
	}

	if p.ctx != nil {
		ctx, cancel := context.WithCancel(p.ctx)
		p.ctx = ctx
		p.cancel = cancel
	}

	p.pool = make(chan *sftp.Client, p.maxSize)
	p.running = make(chan *sftp.Client, p.maxSize)

	return p
}

func (p *WorkerPool) newClient() (*sftp.Client, error) {

	p.mu.Lock()
	defer p.mu.Unlock()

	client, err := sftp.NewClient(p.conn)

	if err == nil {
		p.cls = append(p.cls, client.Close)
	}

	return client, err
}

func (p *WorkerPool) gerWorker() (*sftp.Client, error) {

	availableNew := p.maxSize - (len(p.pool) + len(p.running))

	if availableNew > 0 {

		client, err := p.newClient()

		if err != nil {
			return nil, err
		}
		p.pool <- client
		return client, nil
	}
	w, _ := <-p.pool

	return w, nil

}

func (p *WorkerPool) DownloadFile(src, dest string) {

	file := fileTransfer{
		direction: downloadType,
		src:       src,
		dest:      dest,
	}

	p.transferFile(file)
}

func (p *WorkerPool) UploadFile(src, dest string) {

	file := fileTransfer{
		direction: uploadType,
		src:       src,
		dest:      dest,
	}

	p.transferFile(file)
}

func (p *WorkerPool) transferFile(file fileTransfer) {

	worker, err := p.gerWorker()
	if err != nil {
		p.mu.Lock()
		defer p.mu.Unlock()
		p.err = multierror.Append(p.err, errors.Wrapf(err, "(%s) %s -> %s", file.direction, file.src, file.dest))
		return
	}
	p.wg.Add(1)

	go func() {
		p.running <- worker
		var fn transferFn

		switch file.direction {
		case downloadType:
			fn = p.download
		case uploadType:
			fn = p.upload
		}

		err := transferWithContext(p.ctx, fn, worker, file.src, file.dest)

		if err != nil {
			p.mu.Lock()
			p.err = multierror.Append(p.err, err)
			p.mu.Unlock()
		}
		<-p.running
		p.wg.Done()
	}()

}

func (p *WorkerPool) Wait() {

	p.wg.Wait()

}

func (p *WorkerPool) Close() {

	p.cancel()
	p.mu.Lock()
	defer p.mu.Unlock()

	for _, cl := range p.cls {
		_ = cl() // TODO Сделайть чтение ошибок закрытия
	}

	close(p.pool)
	close(p.running)

}

func (p *WorkerPool) ActiveWorkers() int {
	p.mu.Lock()
	defer p.mu.Unlock()

	return len(p.running)
}
