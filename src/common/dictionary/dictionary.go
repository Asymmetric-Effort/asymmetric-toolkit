package dictionary_old

type Dictionary struct {
	//header
	//content
	//footer
}

type IOMode uint8

const (
	READ   IOMode = 0
	WRITE  IOMode = 1
	CREATE IOMode = 2
)

func (o *Dictionary) Open(fileName string, mode IOMode)

func (o *Dictionary) Reader() string {

}

func (o *Dictionary) Write(s string) {

}
