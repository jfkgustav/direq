// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package musician

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"github.com/jfkgustav/direq/model"
	"github.com/jfkgustav/direq/view"
	"sort"
)

func Index(requestedSongs []model.SongRequest) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)

		sort.Slice(requestedSongs, func(i, j int) bool {
			return requestedSongs[i].TimeInQueue > requestedSongs[j].TimeInQueue
		})
		templ_7745c5c3_Var2 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
			templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
			if !templ_7745c5c3_IsBuffer {
				defer func() {
					templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
					if templ_7745c5c3_Err == nil {
						templ_7745c5c3_Err = templ_7745c5c3_BufErr
					}
				}()
			}
			ctx = templ.InitializeContext(ctx)
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"container px-2 flex flex-col space-y-4 lg:w-1/2 text-xl text-neutral-100\"><h1 class=\"text-center pt-4 text-neutral-100 font-bold text-2xl text-center\">Önskade låtar</h1><div class=\"flex container px-4 space-x-4 justify-center \"><button id=\"queueButton\" class=\"w-1/4 text-center outline outline-indigo-400 py-4 !bg-neutral-200 bg-neutral-500 rounded-md text-black\">Kö </button> <button id=\"popButton\" class=\"w-1/4 text-center outline-indigo-400 py-4 rounded-md bg-neutral-500 text-black\">Populärt </button></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			for _, request := range requestedSongs {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex flex-col space-y-2 p-4 rounded-md container bg-neutral-800\"><a>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				templ_7745c5c3_Err = RequestedSongCard(request).Render(ctx, templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</a></div>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><script>\n\t\t\tconst sorter = document.querySelectorAll('.sorter');\n\t\t\tconst queueButton = document.getElementById('queueButton');\n\t\t\tconst popularityButton = document.getElementById('popButton');\n\n\t\t\tqueueButton.addEventListener('click', () => {\n\t\t\t\t\tif (!queueButton.classList.contains('outline')) {\n\t\t\t\t\t\tsorter.forEach(object => {\n\t\t\t\t\t\t\tobject.classList.toggle('hidden');\n\t\t\t\t\t\t});\n\t\t\t\t\t\tqueueButton.classList.toggle('outline')\n\t\t\t\t\t\tpopularityButton.classList.toggle('outline')\n\t\t\t\t\t\tqueueButton.classList.toggle('!bg-neutral-200')\n\t\t\t\t\t\tpopularityButton.classList.toggle('!bg-neutral-200')\n    \t\t\t}\n\t\t\t});\n\n\t\t\tpopularityButton.addEventListener('click', () => {\n\t\t\t\t\tif (!popularityButton.classList.contains('outline')) {\n\t\t\t\t\t\tsorter.forEach(object => {\n\t\t\t\t\t\t\tobject.classList.toggle('hidden');\n\t\t\t\t\t\t});\n\t\t\t\t\t\tqueueButton.classList.toggle('outline')\n\t\t\t\t\t\tpopularityButton.classList.toggle('outline')\n\t\t\t\t\t\tqueueButton.classList.toggle('!bg-neutral-200')\n\t\t\t\t\t\tpopularityButton.classList.toggle('!bg-neutral-200')\n\t\t\t\t\t}\n\t\t\t});\n\n\t\t</script>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = view.Layout().Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
