<template>
<div>
    <el-row>
        <el-col :span="15">
            <div ref="editor" :id="id" class="editor"></div>
        </el-col>
        <el-col :offset="1" :span="8">
            
            
                <el-col v-for="(item,key) in items" :key="key" :span="24" class="form-item">
                    <el-input v-model="luaForm.keys[key]" placeholder="Please input">
                        <template #prepend>KEYS[{{key+1}}]</template>
                    </el-input>
                    <el-input v-model="luaForm.argv[key]" placeholder="Please input">
                        <template #prepend>ARGV[{{key+1}}]</template>
                    </el-input>
                </el-col>
            
        </el-col>
    </el-row>
    <el-row class="run">
        <el-col>
            <el-button @click="runLua" type="primary">运行</el-button>
        </el-col>
    </el-row>
    <el-row>
        <el-col>
            <el-input :rows="6" v-model="result" disabled type="textarea" />
        </el-col>
    </el-row>
</div>    
</template>

<script>
import * as monaco from 'monaco-editor'
import { language as luaLanguage } from 'monaco-editor/esm/vs/basic-languages/lua/lua.js';
import { reactive, ref, toRaw } from '@vue/reactivity'
import { onMounted } from '@vue/runtime-core'
import { luaRun } from '@/api/index.js'

export default {
    name:"LuaView",
    props: {
        sk:{
            type:String
        },
        db:{
            type:Number
        }
    },
    setup(props) {

        let luaForm = reactive({
            sk:props.sk,
            db:props.db,
            keys:[],
            argv:[],
            script:""
        })

        let result = ref("")
        let id = ref(props.sk+props.db+'lua')

        let items = [1,1,1,1,1];

        const content = ref("")
        const editor = ref(null)
        const initEditor = () => {
        // 初始化编辑器，确保dom已经渲染
        editor.value = monaco.editor.create(document.getElementById(id.value), {
            value: 'return redis.call("GET",KEYS[1])',
            language: 'lua',
            theme: 'vs-dark',
            selectOnLineNumbers: true,
            roundedSelection: false,
            readOnly: false,
            cursorStyle: 'line',
            automaticLayout: true,
            glyphMargin: true,
            useTabStops: false,
            fontSize: 15,
            autoIndent: true,
            quickSuggestionsDelay: 10,
        });
        
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
            luaForm.script = content.value
            luaRun(luaForm).then(res=>{
                result.value = res.data.data
            })
        }

        onMounted(() => {
            initEditor()
        })

        return {
            id,
            items,
            result,
            luaForm,
            editor,
            runLua
        }
    },
}
</script>
<style scoped>
  .editor {
    width:100%;
    height:400px;
  }

  .run{
    margin-top: 10px;
    margin-bottom: 10px;
  }

  .form-item{
    margin-bottom: 10px;
  }
</style>