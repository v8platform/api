package main

import (
	"github.com/Khorevaa/go-v8runner/agent/client/errors"
	"github.com/hashicorp/go-multierror"
	"log"
	"os"
	"path/filepath"
	"sync"
)

func main() {
	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()
	//
	////t := a.T()
	//
	//defer func() {
	//	if err := recover(); err != nil {
	//		fmt.Sprintf(fmt.Sprintf("v8 run agent err:%s", err))
	//	}
	//}()
	//ibPath, _ := ioutil.TempDir("", "1c_DB_")
	//
	//r := runner.NewRunner()
	//ib := infobase.NewFileIB(ibPath)
	//_ = r.Run(ib, infobase.CreateFileInfoBaseOptions{},
	//	runner.WithTimeout(30))
	//
	//_ = r.Run(ib, agent.AgentModeOptions{
	//	Visible:        true,
	//	SSHHostKeyAuto: true,
	//	BaseDir:        "./"},
	//	runner.WithContext(ctx))

	//srcDir := "./agent"
	//destDir := "./src"
	//
	//type file struct {
	//	src string
	//	dest string
	//}
	//
	//fileList := make([]file, 0)
	//e := filepath.Walk(srcDir, func(path string, f os.FileInfo, err error) error {
	//
	//	if f.IsDir() {
	//		return err
	//	}
	//
	//	p, _ := filepath.Rel(srcDir, path)
	//	destFile := filepath.Join(destDir, p)
	//	srcFile, _ := filepath.Abs(path)
	//	fileList = append(fileList, file{srcFile, destFile})
	//	//err = uploadFile(fileTransfer, f.Name(), dest)
	//	return err
	//})
	//
	//println(fileList)
	//println(e)
	treads()
}

func treads() {

	fp := make(filePipe, 5)

	tr := NewTreadsTransfer(5, fp)
	tr.init()

	srcDir := "./agent"
	dest := "./tmp"

	fileList := make([]fileTransfer, 0)

	_ = filepath.Walk(srcDir, func(path string, f os.FileInfo, e error) error {

		if f.IsDir() {
			return e
		}

		p, _ := filepath.Rel(srcDir, path)
		destFile := filepath.Join(dest, p)

		srcFile, _ := filepath.Abs(path)
		ft := fileTransfer{downloadType, srcFile, destFile}
		fileList = append(fileList, ft)

		return e
	})

	tr.wg.Add(len(fileList))

	go func() {
		for _, f := range fileList {
			fp <- f
		}
	}()

	err := tr.Wait()

	if err != nil {
		log.Fatalf("%v", err)
	}
}

type directionType uint

const (
	downloadType directionType = iota
	uploadType
)

type fileTransfer struct {
	direction directionType
	src       string
	dest      string
}
type errorTransfer struct {
	file fileTransfer
	err  error
}

type filePipe chan fileTransfer
type errPipe chan errorTransfer

type transfer struct {

	//agent     *AgentClient
	treads    chan struct{}
	closers   []chan struct{}
	count     int
	filesPipe filePipe
	errsPipe  errPipe
	done      chan bool
	close     chan bool
	errs      []errorTransfer
	wg        *sync.WaitGroup
}

func NewTreadsTransfer(count int, filesPipe filePipe) *transfer {

	done := make(chan bool, 1)

	return &transfer{
		count:     count,
		filesPipe: filesPipe,
		errsPipe:  make(errPipe, count*10),
		done:      done,
		close:     make(chan bool, 1),
		errs:      []errorTransfer{},
		treads:    make(chan struct{}, count),
		wg:        new(sync.WaitGroup),
	}

}

func (t *transfer) init() {

	for i := 0; i < t.count; i++ {

		err, closer := t.treadTransfer(t.filesPipe, t.errsPipe, t.done, t.wg)

		t.closers = append(t.closers, closer)

		if err != nil {
			log.Printf("err create tread transfer %v", err)
		}

	}

	go func() {
		for {
			select {
			case err, ok := <-t.errsPipe:
				if ok {
					t.errs = append(t.errs, err)
				}

			case <-t.close:
				t.Close()
				break
			}
		}

	}()

}

func (t *transfer) Close() {

	for _, tread := range t.closers {
		tread <- struct{}{}
		close(tread)
	}

	t.close <- true

}

func (t *transfer) Wait() error {

	t.wg.Wait()

	t.Close()

	result := new(multierror.Error)

	for _, err := range t.errs {

		result = multierror.Append(result, err.err)
	}

	return result.ErrorOrNil()
}

func (c *transfer) treadTransfer(fp filePipe, errP errPipe, done chan bool, wg *sync.WaitGroup) (error, chan struct{}) {

	//conn, err := c.agent.newConnection()
	//if err != nil {
	//	return err, nil
	//}
	//
	//client, err := sftp.NewClient(conn)
	//
	//if err != nil {
	//	return err, nil
	//}
	//defer client.Close()

	closer := make(chan struct{})

	raiseE := false

	go func() {
		for {

			select {

			case ft, ok := <-fp:

				if !ok {
					continue
				}

				switch ft.direction {

				case downloadType:

					e := downloadFile(ft.src, ft.dest)

					if e != nil || raiseE {
						errP <- errorTransfer{
							file: ft,
							err:  errors.Wrapf(e, "raised"),
						}
					}

					if raiseE {
						raiseE = !raiseE
					}

				case uploadType:
					e := uploadFile(ft.src, ft.dest)

					if e != nil {
						errP <- errorTransfer{
							file: ft,
							err:  e,
						}
					}

				}
				wg.Done()

			case <-done:
				//conn.Close()
				//client.Close()
				break
			case <-closer:
				//conn.Close()
				//client.Close()
				break
			}
		}

	}()

	return nil, closer

}

func uploadFile(src, dest string) error {

	targetDir := filepath.Dir(dest)
	targetDir = filepath.ToSlash(targetDir)
	log.Printf("upload file %s -> %s", src, dest)

	//size := int64(0)

	//err := client.MkdirAll(targetDir)
	//
	//if err != nil {
	//	log.Fatal(err)
	//}

	//// create destination file
	//dstFile, err := client.Create(dest)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer dstFile.Close()
	//
	//// create source file
	//srcFile, err := os.Open(src)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer srcFile.Close()
	//
	//// copy source file to destination file
	//_, err = io.Copy(dstFile, srcFile)
	//if err != nil {
	//	log.Fatal(err)
	//}

	return nil
}

func downloadFile(src, dest string) error {

	targetDir := filepath.Dir(dest)
	targetDir = filepath.ToSlash(targetDir)

	log.Printf("download file %s -> %s", src, dest)

	//size := int64(0)

	//err := os.MkdirAll(targetDir, os.ModePerm)
	//
	//if err != nil {
	//	log.Fatal(err)
	//}

	//// create destination file
	//dstFile, err := os.Create(dest)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer dstFile.Close()
	//
	//// create source file
	//srcFile, err := client.Open(src)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer srcFile.Close()
	//
	//// copy source file to destination file
	//_, err = io.Copy(dstFile, srcFile)
	//if err != nil {
	//	log.Fatal(err)
	//}

	return nil
}
