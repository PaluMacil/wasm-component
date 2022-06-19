package blog

import (
	"github.com/PaluMacil/wasm-component/example/web/blog/article"
)

// BlogState hold state about the BlogComponent
type BlogState struct {
	Articles []article.ArticleState
}

// TODO: a component needs to Render and it needs
// a reference to its "new" constructor for IoC
// also needs hooking
// BlogComponent holds the
type BlogComponent struct {
	//State wc.State[BlogState]
}
