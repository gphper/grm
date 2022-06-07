import { Terminal } from "xterm";
import { FitAddon } from "xterm-addon-fit";
import 'xterm/css/xterm.css'

const term = new Terminal({
    rendererType: "canvas", //渲染类型
    rows: 20, //行数
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
    bellStyle: 'sound',
});

term.writeln("\x1b[38:5:29m连接中...\x1b[0m");
term.write('\x1b[32m链接名称:0> \x1b[0m');

let curr_line = '';
let history_lines = [];
let last = 0
let current = 0;
term.onData((data)=>{
    let code = data.charCodeAt(0);
    switch(code){
        case 127: //backspace
        console.log(curr_line)
            if(curr_line.length > 0){
                curr_line = curr_line.slice(0, -1);
                term.write('\x1b[2K\r');
                term.write('\x1b[32m链接名称:0> \x1b[0m'+curr_line);
            }else{
                term.write("\x07");
            }
            break;
        case 13: //enter
            if (curr_line === 'clear') {
                term.clear()
            }
            if(curr_line != ""){
                history_lines.push(curr_line);
                last = history_lines.length - 1;
                current = last+1;
            }
            term.write('\r\n\x1b[32m链接名称:0> \x1b[0m');
            curr_line = ""
            break;
        case 27: //方向键
            if(data === '\u001b[A'){//up
                if (last >= 0 && current > 0) {
                    current--
                    //当前行有数据的时候进行删除掉在进行渲染上存储的历史数据
                    if(curr_line.length > 0){
                        term.write('\x1b[2K\r');
                        term.write('\x1b[32m链接名称:0> \x1b[0m');
                    }

                    let text = history_lines[current]
                    term.write(text)
                    //重点，一定要记住存储当前行命令保证下次up或down时不会光标错乱覆盖终端提示符
                    curr_line = text
                }else{
                    term.write("\x07");
                }
                console.log('curent'+current);
            }else if(data === '\u001b[B'){//down
                if (last >= 0 && current < last) {
                    current++
                    if(curr_line.length > 0){
                        term.write('\x1b[2K\r');
                        term.write('\x1b[32m链接名称:0> \x1b[0m');
                    }

                    let text = history_lines[current]
                    term.write(text)
                    curr_line = text
                }else{
                    term.write("\x07"); 
                }
                console.log('curent'+current);
            }else{
                term.write("\x07");
            }
            break;            
        default:
            curr_line += data;
            term.write(data);
            break;
    }

    console.log(curr_line)
});


const fitAddon = new FitAddon()

export {term,fitAddon}