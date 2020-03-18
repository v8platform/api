package sshclient

import (
	"github.com/pkg/sftp"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

)

type directionType uint
const (
	downloadType directionType = iota
	uploadType
)
type fileTransfer struct {
	direction directionType
	src  	  string
	dest      string
}
type errorTransfer struct {
	file fileTransfer
	err  error
}

type filePipe chan fileTransfer
type errPipe chan errorTransfer

type transfer struct {

	agent     *AgentClient
	treads 	  []chan bool
	count     int
	filesPipe filePipe
	errsPipe  errPipe
	done      chan bool
	close     chan bool
	errs      []errorTransfer
}

func NewTreadsTransfer(count int, filesPipe filePipe) *transfer {

	return &transfer{
		count:     count,
		treads:    []chan bool,
		filesPipe: filesPipe,
		errsPipe:  make(errPipe, count*10),
		done:      make(chan bool, 1),
		close:     make(chan bool, 1),
		errs:      []errorTransfer{},
	}

}

func (t *transfer) init() {

	for i := 1; t.count == i ; i++ {

		err, closer := t.treadTransfer(t.filesPipe, t.errsPipe, t.done)

		t.treads = append(t.treads, closer)

		if err != nil {
			log.Printf("err create tread transfer %v", err)
		}

	}

	go func() {

		select {
		case <-t.close:
			t.Close()
			break
		case err, ok := <-t.errsPipe:
			if ok {
				t.errs = append(t.errs, err)
			}
		}
	}()

}
func (t *transfer) Close() {

	for _, tread := range t.treads {
		tread <- true
		close(tread)
	}

}


func (c *transfer) treadTransfer(fp filePipe, errP errPipe, done chan bool) (error, chan bool) {

	conn, err := c.agent.newConnection()
	if err != nil {
		return err, nil
	}

	client, err := sftp.NewClient(conn)

	if err != nil {
		return err, nil
	}
	defer client.Close()

	closer := make(chan bool)

	go func() {
		for {

			select {

			case ft, ok := <-fp:
				
				if !ok {
					continue	
				}
								
				switch ft.direction {

				case downloadType:

					e := downloadFile(client, ft.src, ft.dest)

					if e != nil {
						errP <- errorTransfer{
							file: ft,
							err:  e,
						}
					}
					
				case uploadType:
					e := uploadFile(client, ft.src, ft.dest)

					if e != nil {
						errP <- errorTransfer{
							file: ft,
							err:  e,
						}
					}
				}

			case <- done:
				conn.Close()
				client.Close()
				break
			case <- closer:
				conn.Close()
				client.Close()
				break
			}
		}

	}()

	return nil, closer

}


func uploadFile(client *sftp.Client, src, dest string) error {

	targetDir := filepath.Dir(dest)
	targetDir = filepath.ToSlash(targetDir)

	//size := int64(0)

	err := client.MkdirAll(targetDir)

	if err != nil {
		log.Fatal(err)
	}

	// create destination file
	dstFile, err := client.Create(dest)
	if err != nil {
		log.Fatal(err)
	}
	defer dstFile.Close()

	// create source file
	srcFile, err := os.Open(src)
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

func downloadFile(client *sftp.Client, src, dest string) error {

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

func (c *AgentClient) CopyDirToWithTreads(src, dest string, treadCount int) error {

	conn, err := c.newConnection()
	if err != nil {
		return err
	}

	t, err :=

	if err != nil {
		return err
	}
	defer fileTransfer.Close()

	srcDir := filepath.Dir(src)
	srcDir = filepath.ToSlash(srcDir)


	go func() {
		fileList := make([]fileTransfer, 0)
		err = filepath.Walk(srcDir, func(path string, f os.FileInfo, e error) error {

			if f.IsDir() {
				return e
			}

			p, _ := filepath.Rel(srcDir, path)
			destFile := filepath.Join(dest, p)

			srcFile, _ := filepath.Abs(path)
			fileList = append(fileList, fileTransfer{downloadType, srcFile, destFile})
			return e
		})
	}()


	if err != nil {
		return err
	}

	for _, f := range fileList {

		err := uploadFile(fileTransfer, f.src, f.dest)
		if err != nil {
			return err
		}

	}

	err = conn.Close()

	return err

}


func (c *AgentClient) CopyDirTo(src, dest string) error {

	conn, err := c.newConnection()
	if err != nil {
		return err
	}

	fileTransfer, err := sftp.NewClient(conn)

	if err != nil {
		return err
	}
	defer fileTransfer.Close()

	srcDir := filepath.Dir(src)
	srcDir = filepath.ToSlash(srcDir)



	fileList := make([]file, 0)
	err = filepath.Walk(srcDir, func(path string, f os.FileInfo, e error) error {

		if f.IsDir() {
			return e
		}

		p, _ := filepath.Rel(srcDir, path)
		destFile := filepath.Join(dest, p)

		srcFile, _ := filepath.Abs(path)
		fileList = append(fileList, file{srcFile, destFile})
		return e
	})

	if err != nil {
		return err
	}

	for _, f := range fileList {

		err := uploadFile(fileTransfer, f.src, f.dest)
		if err != nil {
			return err
		}

	}

	err = conn.Close()

	return err

}

func (c *AgentClient) CopyFileTo(src, dest string) error {

	conn, err := c.newConnection()
	if err != nil {
		return err
	}

	fileTransfer, err := sftp.NewClient(conn)

	if err != nil {
		return err
	}
	defer fileTransfer.Close()

	err = uploadFile(fileTransfer, src, dest)
	if err != nil {
		return err
	}

	err = conn.Close()

	return err

}

func downloadWithWalker(client *sftp.Client, w, src, dest string, download func(srcFile, destFile string) error) error {

	walker := client.Walk(w)
	//walker.SkipDir()
	for walker.Step() {
		if err := walker.Err(); err != nil {
			continue
		}

		stat := walker.Stat()
		log.Printf("file %s", walker.Path())
		if stat.IsDir() && w == walker.Path() {
			log.Printf("skip dir %s", walker.Path())
			continue
		}

		if stat.IsDir() {
			log.Printf("go to dir %s", walker.Path())
			err := downloadWithWalker(client, walker.Path(), src, dest, download)

			if err != nil {
				return err
			}
		}

		path := walker.Path()

		p, _ := filepath.Rel(src, path)
		destFile := filepath.Join(dest, p)

		err := download(path, destFile)

		if err != nil {
			return err
		}
	}
	return nil
}

func (c *AgentClient) CopyDirFrom(src, dest string) error {

	conn, err := c.newConnection()
	if err != nil {
		return err
	}

	fileTransfer, err := sftp.NewClient(conn)

	if err != nil {
		return err
	}
	defer fileTransfer.Close()

	downloadFn := func(srcFile, destFile string) error {
		e := downloadFile(fileTransfer, srcFile, destFile)
		if e != nil {
			return e
		}
		return nil
	}

	err = downloadWithWalker(fileTransfer, src, src, dest, downloadFn)
	if err != nil {
		return err
	}

	err = conn.Close()

	return err
}

func (c *AgentClient) CopyFileFrom(src, dest string) error {

	conn, err := c.newConnection()
	if err != nil {
		return err
	}

	fileTransfer, err := sftp.NewClient(conn)

	if err != nil {
		return err
	}
	defer fileTransfer.Close()

	if strings.HasSuffix(dest, "/") {
		dest = filepath.Join(dest, filepath.Base(src))
	}

	err = downloadFile(fileTransfer, src, dest)
	if err != nil {
		return err
	}

	err = conn.Close()

	return err
}
