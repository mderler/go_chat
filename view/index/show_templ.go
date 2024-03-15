// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package index

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"github.com/mderler/go_chat/model"
	"github.com/mderler/go_chat/view/layout"
	"strings"
)

func extractFirstCharacters(name string) string {
	words := strings.Split(name, " ")
	wordCount := len(words)
	if wordCount > 1 {
		return strings.ToUpper(words[0][:1] + words[wordCount-1][:1])
	}
	return strings.ToUpper(name[:2])
}

func userIcon(name string, color string) templ.Component {
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
		var templ_7745c5c3_Var2 = []any{"h-8 w-8 mr-2 border rounded-full text-xl text-center", bgColor(color)}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var2...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<span class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ.CSSClasses(templ_7745c5c3_Var2).String()))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(extractFirstCharacters(name))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/index/show.templ`, Line: 21, Col: 32}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func Show(user model.GetUserForChatByIdRow, chatParams ChatParams, contactedUsers []model.GetContactedUsersRow) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var4 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var4 == nil {
			templ_7745c5c3_Var4 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var5 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
			if !templ_7745c5c3_IsBuffer {
				templ_7745c5c3_Buffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"bg-gray-100 flex h-screen text-white bg-gradient-to-b from-slate-600 to-[#240000]\"><div class=\"flex flex-col border-r pt-2\"><div class=\"flex flex-row gap-2 items-center select-none\"><div class=\"text-xl ml-2 flex items-center\"><div class=\"ml-2 flex items-center\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = userIcon(user.FullName, user.Color).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<span class=\"select-text\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var6 string
			templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(user.FullName)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/index/show.templ`, Line: 35, Col: 48}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span></div><button hx-post=\"/logout\">&#x1F6AA;</button></div><button id=\"new-chat-button\" class=\"ml-auto mr-2\">&#x270F;</button></div><section class=\"w-80 grow\"><ul class=\"flex flex-col gap-2 mx-2 mt-2\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = ShowContactedUsers(contactedUsers).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</ul></section></div><div id=\"chat-container\" class=\"flex flex-col grow bg-slate-950/30 justify-center h-screen max-h-screen\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if len(chatParams.Messages) > 0 {
				templ_7745c5c3_Err = ShowChat(chatParams).Render(ctx, templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex flex-col text-pretty border-[16px] border-dashed border-gray-300 rounded-xl h-72 max-w-92 p-8 self-center\"><div class=\"text-4xl text-gray-300 font-extrabold\">Here will be your chats!</div><div class=\"text-xl italic text-gray-400 text-center text-pretty w-64 m-auto\">Click on the \"New Chat\" Icon at the top left to start a conversation.</div></div>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div><dialog id=\"new-chat-dialog\" class=\"fixed text-white border border-black rounded-xl px-4 py-4 bg-slate-900 w-[48rem]\"><form method=\"dialog\"><div class=\"flex flex-row justify-between pb-2 mb-2\"><div class=\"cursor-pointer\"><span class=\"text-xl font-bold border-b-2 border-teal-400\" hx-get=\"/new-chat\" hx-target=\"#new-chat-or-group-content\">New Chat</span> <span class=\"text-xl font-bold border-b-2 border-slate-700\" hx-get=\"/new-group\" hx-target=\"#new-chat-or-group-content\">New Group</span></div><button type=\"cancel\" class=\"text-4xl text-red-500 select-none\">&times;</button></div><div id=\"new-chat-or-group-content\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = ShowNewChat().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></form></dialog><script>\n\t\t\tvar messageContainer\n\n\t\t\tconst socket = new WebSocket(`ws://${window.location.host}/ws`)\t\n\t\t\tsocket.binaryType = 'arraybuffer'\n\n\t\t\tsocket.onmessage = (event) => {\n\t\t\t\tconst message = unpackMessage(event.data)\n\n\t\t\t\tif (!messageContainer) {\n\t\t\t\t\treturn\n\t\t\t\t}\n\n\t\t\t\tmessageContainer.innerHTML += message.text\n\t\t\t\tif (messageContainer.scrollHeight - messageContainer.scrollTop < 1000) {\n\t\t\t\t\tmessageContainer.scrollTop = messageContainer.scrollHeight\n\t\t\t\t}\n\t\t\t}\n\n\t\t\tsocket.onclose = (event) => {\n\t\t\t\tconsole.log('Socket closed', event)\n\t\t\t}\n\n\t\t\tsocket.onerror = (event) => {\n\t\t\t\tconsole.error('Socket error', event)\n\t\t\t}\n\n\t\t\tfunction unpackMessage(data) {\n\t\t\t\tconst isGroup = data[0] !== 0;\n\n\t\t\t\tconst userOrGroupId = new DataView(data).getBigUint64(1, true);\n\n\t\t\t\tconst textBytes = data.slice(9);\n\t\t\t\tconst text = new TextDecoder().decode(textBytes);\n\n\t\t\t\treturn { isGroup, userOrGroupId, text };\n\t\t\t}\n\n\t\t\tfunction sendMessage(isGroup, userOrGroupId, text) {\n\t\t\t\tconst typeBit = isGroup ? 1 : 0;\n\n\t\t\t\tconst textBytes = new TextEncoder().encode(text);\n\n\t\t\t\tconst buffer = new ArrayBuffer(9 + textBytes.length);\n\t\t\t\tconst view = new DataView(buffer);\n\n\t\t\t\tview.setUint8(0, typeBit);\n\n\t\t\t\tview.setBigUint64(1, BigInt(userOrGroupId), true);\n\n\t\t\t\tfor (let i = 0; i < textBytes.length; i++) {\n\t\t\t\t\tview.setUint8(i + 9, textBytes[i]);\n\t\t\t\t}\n\n\t\t\t\tsocket.send(buffer)\n\t\t\t}\n\n\t\t\tconst dialog = document.getElementById('new-chat-dialog')\n\n\t\t\tdocument.getElementById('new-chat-button')\n\t\t\t\t.addEventListener('click', () => {\n\t\t\t\t\tdialog.showModal()\n\t\t\t\t})\n\t\t\t\n\t\t\tconst userSearch = document.getElementById('user-search')\n\n\t\t\tfunction selectContact() {\n\t\t\t\tdialog.close()\n\t\t\t\tuserSearch.value = ''\n\t\t\t}\n\t\t</script>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if !templ_7745c5c3_IsBuffer {
				_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = layout.Base().Render(templ.WithChildren(ctx, templ_7745c5c3_Var5), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
