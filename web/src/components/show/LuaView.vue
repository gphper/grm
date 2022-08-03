<template>
<div>
    <el-row>
        <el-col :span="15">
            <div ref="editor" id="codeEditBox"></div>
        </el-col>
        <el-col :offset="1" :span="8">
            <el-input v-model="input1" placeholder="Please input">
                <template #prepend>KEYS[1]</template>
            </el-input>
            <el-input v-model="input1" placeholder="Please input">
                <template #prepend>ARGV[1]</template>
            </el-input>
        </el-col>
    </el-row>
    <el-row class="run">
        <el-col>
            <el-button @click="runLua" type="primary">运行</el-button>
        </el-col>
    </el-row>
    <el-row>
        <el-col>
            <el-input disabled type="textarea" />
        </el-col>
    </el-row>
</div>    
</template>

<script>
import * as monaco from 'monaco-editor'
import { language as luaLanguage } from 'monaco-editor/esm/vs/basic-languages/lua/lua.js';
import { ref, toRaw } from '@vue/reactivity'
import { onMounted } from '@vue/runtime-core'
export default {
    name:"LuaView",
    setup() {
        const content = ref("")
        const editor = ref(null)
        const initEditor = () => {
        // 初始化编辑器，确保dom已经渲染
        editor.value = monaco.editor.create(document.getElementById('codeEditBox'), {
            value: 'please write lua script', //编辑器初始显示文字
            language: 'lua', //此处使用的python，其他语言支持自行查阅demo
            theme: 'vs-dark', //官方自带三种主题vs, hc-black, or vs-dark
            selectOnLineNumbers: true,//显示行号
            roundedSelection: false,
            readOnly: false, // 只读
            cursorStyle: 'line', //光标样式
            automaticLayout: true, //自动布局
            glyphMargin: true, //字形边缘
            useTabStops: false,
            fontSize: 15, //字体大小
            autoIndent: true, //自动布局
            quickSuggestionsDelay: 10, //代码提示延时
        });
        // 监听值的变化
        // editor.value.onDidChangeModelContent((val) => {
        //     console.log(val.changes[0].text)
        // })

        monaco.languages.registerCompletionItemProvider('lua', {
                    provideCompletionItems: function () {
                        let suggestions = [];
                        luaLanguage.keywords.forEach(item => {
                            suggestions.push({
                                label: item,
                                kind: monaco.languages.CompletionItemKind.Keyword,
                                insertText: item
                            });
                        })
                        return {
                            // 最后要返回一个数组
                            suggestions:suggestions
                        };
                    },
                });

        };
        
        const runLua = ()=>{
            content.value = toRaw(editor.value).getValue()
            console.log(content.value)
        }

        onMounted(() => {
            initEditor()
        })

        return {
            editor,
            runLua
        }
    },
}
</script>
<style scoped>
  #codeEditBox {
    width:100%;
    height:400px;
  }

  .run{
    margin-top: 10px;
    margin-bottom: 10px;
  }
</style>