import { Terminal } from "xterm";
import { FitAddon } from "xterm-addon-fit";
import 'xterm/css/xterm.css'

const term = new Terminal({
    rendererType: "canvas", //渲染类型
    rows: 20, //行数
    // cols: 40, // 不指定行数，自动回车后光标从下一行开始
    // convertEol: true, //启用时，光标将设置为下一行的开头
    scrollback: 50, //终端中的回滚量
    disableStdin: false, //是否应禁用输入
    windowsMode: true, // 根据窗口换行
    cursorStyle: "underline", //光标样式
    cursorBlink: true, //光标闪烁
    theme: {
      foreground: "#ECECEC", //字体
      background: "#000000", //背景色
      cursor: "help", //设置光标
      lineHeight: 20,
    },
});
term.writeln("连接中...");

term.write('>');
term.onKey((e)=>{
    console.log(e.key.charCodeAt(0))
    let code = e.key.charCodeAt(0);
    switch(code){
        case 127: //backspace
            term.write("\b \b");
            break;
        case 13: //enter
            term.writeln('');
            term.write('>');
            break;    
        default:
            term.write(e.key)
            break;    
    }
})


const fitAddon = new FitAddon()

export {term,fitAddon}