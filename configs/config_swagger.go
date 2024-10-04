package configs

type ISwagger interface {
	Next() *bool
	BasePath() string
	FilePathfilePath() string
	Path() string
	Title() string
	CacheAge() int
}

type swagger struct {
	next     *bool
	basePath string
	filePath string
	path     string
	title    string
	cacheAge int
}

func (c *config) Swagger() ISwagger {
	return c.swagger
}

func (s *swagger) Next() *bool              { return s.next }
func (s *swagger) BasePath() string         { return s.basePath }
func (s *swagger) FilePathfilePath() string { return s.filePath }
func (s *swagger) Path() string             { return s.path }
func (s *swagger) Title() string            { return s.title }
func (s *swagger) CacheAge() int            { return s.cacheAge }
