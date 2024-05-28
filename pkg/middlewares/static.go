package middlewares

import (
	"embed"
	"ginfx-template/ui"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
	"strings"
)

const INDEX = "index.html"

type ServeFileSystem interface {
	http.FileSystem
	Exists(prefix string, path string) bool
}

type EmbedFileSystem struct {
	http.FileSystem
	root    string
	indexes bool
}

func (em *EmbedFileSystem) Open(name string) (http.File, error) {
	return em.FileSystem.Open(path.Join(em.root, name))
}

func EmbedFile(root string, fs *embed.FS, indexes bool) *EmbedFileSystem {
	return &EmbedFileSystem{
		FileSystem: http.FS(fs),
		root:       root,
		indexes:    indexes,
	}
}

func (em *EmbedFileSystem) Exists(prefix string, filepath string) bool {
	if p := strings.TrimPrefix(filepath, prefix); len(p) < len(filepath) {
		file, err := em.Open(p)
		if err != nil {
			return false
		}
		if stats, err := file.Stat(); err != nil {
			return false
		} else {
			if !em.indexes {
				if stats.IsDir() {
					index := path.Join(p, INDEX)
					if indexFile, err := em.Open(index); err != nil {
						return false
					} else if _, err = indexFile.Stat(); err != nil {
						return false
					}
				}
			}
			return true
		}
	}
	return false
}

func ServeRoot(urlPrefix, root string) gin.HandlerFunc {
	return Serve(urlPrefix, EmbedFile(root, &ui.RESOURCE, false))
}

// Serve Static returns a middleware handler that serves static files in the given directory.
func Serve(urlPrefix string, fs ServeFileSystem) gin.HandlerFunc {
	fileServer := http.FileServer(fs)
	if urlPrefix != "" {
		fileServer = http.StripPrefix(urlPrefix, fileServer)
	}
	return func(c *gin.Context) {
		if fs.Exists(urlPrefix, c.Request.URL.Path) {
			fileServer.ServeHTTP(c.Writer, c.Request)
			c.Abort()
		}
	}
}
