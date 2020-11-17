package source

import (
	"testing"
)

func TestSourceGenerateDictionary(t *testing.T) {
	t.SkipNow()
	/*
	dictionary:=func() string {
		/*
			Setup the test by creating a test dictionary file.
		*/
	/*
		checkError := func(err error) {
			if err != nil {
				panic(err)
			}
		}

		var file *os.File = func() (fh *os.File) {
			var err error
			tempDir := os.TempDir()
			if !fileUtils.DirExists(tempDir) {
				checkError(os.MkdirAll(tempDir, os.ModePerm))
			}
			filename := filepath.Join(tempDir, fmt.Sprintf("testDictionary.%d.txt", int(time.Now().Unix())))
			fh, err = os.Create(filename)
			checkError(err)
			return fh
		}()
		var writer writer2.Writer
		//writeDict:=writer.Setup(file)
		defer writer.Close()
		for i := 1; i < 10; i++ {
			//writeDict(fmt.Sprintf("Password%d", i))
		}
		//We have created the dictionary and we can now close the file and return the filename
		return file.Name()
	}()
	/*
		Setup the test by configuring our dictionary source.
	 */

	/*
		Run the test.  Consume the dictionary into the FIFO queue then verify it all worked.
	 */
}
