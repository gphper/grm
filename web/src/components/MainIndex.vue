<template>
<div class="box" ref="box">
    <div class="left" ref="left">
    <MenuSection></MenuSection>
    </div>
    <div class="resize" ref="resize" @mousedown="resizeMouse($event)" title="收缩侧边栏">
    ⋮
    </div>
    <div class="mid" ref="mid">
    <MainSection :tagData="tabs"></MainSection>
    </div>

    <DataForm></DataForm>
    <SettingFrom></SettingFrom>
</div>
</template>

<script>
import { reactive, ref } from '@vue/reactivity'
import MainSection from '@/components/MainSection.vue'
import MenuSection from '@/components/MenuSection.vue'
import DataForm from "@/components/index/DataForm.vue"
import SettingFrom from "@/components/index/SettingForm.vue"

export default {
    name:"MainIndex",
    components:{
        MainSection,
        MenuSection,
        DataForm,
        SettingFrom
    },
    setup() {

    let left = ref(null);
    let mid = ref(null);
    let box = ref(null);
    let resize = ref(null);
    let title = ref("");

    let tabs = reactive([
      {
        title: 'Tab 1',
        name: '1',
        content: 'Tab 1 content',
      },
      {
        title: 'Tab 2',
        name: '2',
        content: 'Tab 2 content',
      },
    ]);
    
    const resizeMouse = ()=>{
      document.onmousemove = function(e){
        if(e.clientX > 400 && (box.value.clientWidth - e.clientX)>300){
          //左边容器
          left.value.style.width=(e.clientX-10) +'px'
          mid.value.style.width=(box.value.clientWidth - e.clientX)+'px';
          resize.value.style.color='#337ecc';
        }
      }

      document.onmouseup = function(){
          document.onmousemove = null;
          document.onmouseup = null;
          resize.value.style.color='white';
      }
    }

    return{
      left,
      mid,
      box,
      resize,
      tabs,
      title,
      resizeMouse,
    }
  },
}
</script>

<style scoped>
.box {
    width: 100%;
    height: 100%;
    margin: 0 0px;
    overflow: hidden;
    box-shadow: -1px 9px 10px 3px rgba(0, 0, 0, 0.11);
    background-color:    #ecf5ff;
}

.left {
    width: calc(32% - 10px);
    height: 100%;
    background: #FFFFFF;
    float: left;
}

.resize {
    cursor: col-resize;
    float: left;
    position: relative;
    top: 45%;
    background-color:  #79bbff;
    border-radius: 5px;
    margin-top: -10px;
    width: 10px;
    height: 50px;
    background-size: cover;
    background-position: center;
    font-size: 32px;
    color: white;
    -moz-user-select: none;
    -webkit-user-select: none;
    -ms-user-select: none;
    -khtml-user-select: none;
    user-select: none;
}

.resize:hover {
    color:  #337ecc;
}
.mid {
    float: left;
    width: 68%;
    height: 100%;
    background: #fff;
    box-shadow: -1px 4px 5px 3px rgba(0, 0, 0, 0.11);
}
</style>