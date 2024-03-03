// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package index

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "github.com/mderler/go_chat/view/layout"

func Show() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var2 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
			if !templ_7745c5c3_IsBuffer {
				templ_7745c5c3_Buffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<body class=\"base-grid bg-gray-100 dark:bg-slate-900 text-black dark:text-white\"><div class=\"border-r ml-2 flex items-center justify-around gap-2\"><div class=\"text-xl\">Matthias Derler</div><div><label for=\"darkMode\">Dark Mode:</label> <input id=\"darkMode\" type=\"checkbox\" onchange=\"onDarkModeChange\"></div></div><div class=\"ml-2 text-2xl flex items-center\"><div class=\"h-8 w-8 mr-2 border rounded-full bg-blue-400 text-xl text-center\">SD</div><span>Stefan Derler</span></div><section class=\"dark:bg-slate-600 w-80 border-r\"><ul class=\"flex flex-col gap-2 mx-2 mt-2\"><div class=\"border border-black rounded-md flex px-1 py-1\"><div class=\"h-8 w-8 mr-2 border rounded-full bg-blue-400 text-xl text-center\">SD</div><span>Stefan Derler</span></div></ul></section><section class=\"dark:bg-slate-600 chat-grid h-full border-inherit\"><div class=\"border-inherit flex flex-col overflow-y-auto\"><div class=\"left-chat\"><div class=\"font-bold text-green-300\">Stefan Derler</div><p class=\"dark:text-white\">Lorem ipsum dolor sit amet consectetur adipisicing elit. Maiores veritatis quibusdam fugit odio laborum quaerat iure rem nisi quidem quasi itaque ut, harum commodi ea in distinctio quod dolor placeat.</p></div><div class=\"left-chat\"><div class=\"font-bold text-green-300\">Stefan Derler</div><p class=\"dark:text-white\">Lorem ipsum dolor sit amet consectetur adipisicing elit. Maiores veritatis quibusdam fugit odio laborum quaerat iure rem nisi quidem quasi itaque ut, harum commodi ea in distinctio quod dolor placeat.</p></div><div class=\"left-chat\"><div class=\"font-bold text-green-300\">Stefan Derler</div><p class=\"dark:text-white\">Lorem ipsum dolor sit amet consectetur adipisicing elit. Maiores veritatis quibusdam fugit odio laborum quaerat iure rem nisi quidem quasi itaque ut, harum commodi ea in distinctio quod dolor placeat.</p></div><div class=\"left-chat\"><div class=\"font-bold text-green-300\">Stefan Derler</div><p class=\"dark:text-white\">Lorem ipsum dolor sit amet consectetur adipisicing elit. Maiores veritatis quibusdam fugit odio laborum quaerat iure rem nisi quidem quasi itaque ut, harum commodi ea in distinctio quod dolor placeat.</p></div><div class=\"left-chat\"><div class=\"font-bold text-green-300\">Stefan Derler</div><p class=\"dark:text-white\">Lorem ipsum dolor sit amet consectetur adipisicing elit. Maiores veritatis quibusdam fugit odio laborum quaerat iure rem nisi quidem quasi itaque ut, harum commodi ea in distinctio quod dolor placeat.</p></div><div class=\"left-chat\"><div class=\"font-bold text-green-300\">Stefan Derler</div><p class=\"dark:text-white\">Lorem ipsum dolor sit amet consectetur adipisicing elit. Maiores veritatis quibusdam fugit odio laborum quaerat iure rem nisi quidem quasi itaque ut, harum commodi ea in distinctio quod dolor placeat.</p></div><div class=\"left-chat\"><div class=\"font-bold text-green-300\">Stefan Derler</div><p class=\"dark:text-white\">Lorem ipsum dolor sit amet consectetur adipisicing elit. Maiores veritatis quibusdam fugit odio laborum quaerat iure rem nisi quidem quasi itaque ut, harum commodi ea in distinctio quod dolor placeat.</p></div><div class=\"left-chat\"><div class=\"font-bold text-green-300\">Stefan Derler</div><p class=\"dark:text-white\">Lorem ipsum dolor sit amet consectetur adipisicing elit. Maiores veritatis quibusdam fugit odio laborum quaerat iure rem nisi quidem quasi itaque ut, harum commodi ea in distinctio quod dolor placeat.</p></div><div class=\"left-chat\"><div class=\"font-bold text-green-300\">Stefan Derler</div><p class=\"dark:text-white\">Lorem ipsum dolor sit amet consectetur adipisicing elit. Maiores veritatis quibusdam fugit odio laborum quaerat iure rem nisi quidem quasi itaque ut, harum commodi ea in distinctio quod dolor placeat.</p></div><div class=\"right-chat\"><div class=\"font-bold text-pink-400\">Matthias Derler</div><p class=\"dark:text-white\">Lorem ipsum dolor sit amet consectetur adipisicing elit. Maiores veritatis quibusdam fugit odio laborum quaerat iure rem nisi quidem quasi itaque ut, harum commodi ea in distinctio quod dolor placeat.</p></div><div class=\"right-chat mt-[0.1rem]\"><p class=\"dark:text-white\">Lorem ipsum elit.</p></div></div><div class=\"mt-auto mx-5 mb-3 shadow-lg bg-slate-900 rounded-md px-2 pt-2 pb-1 flex flex-col align-middle\"><textarea class=\"w-full bg-slate-900 text-white chat-input outline-none\" placeholder=\"Type a message...\"></textarea></div></section></body>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if !templ_7745c5c3_IsBuffer {
				_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = layout.Base().Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
