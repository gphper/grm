<template>
    <el-drawer
      v-model="drawerAble"
      :title="currentTitle"
      @close="closeParent"
      >
        <span v-html="currentInfo"></span>
    </el-drawer>
</template>

<script>
import { ref, watch } from '@vue/runtime-core';
export default {
    name:"ConnInfo",
    props:{
        drawer:{
            type:Boolean
        },
        title:{
            type:String
        },
        info:{
            type:String
        }
    },
    setup(props,context) {
        let drawerAble = ref(props.drawer);
        let currentTitle = ref(props.title);
        let currentInfo = ref(props.info);

        const closeParent = ()=>{
            context.emit("close");
        }
        
        watch(
            ()=>[props.drawer,props.title,props.info],
            (value) => {
                currentTitle.value = value[1];
                drawerAble.value = value[0];
                currentInfo.value = value[2]
            }
        );

        return {
            currentTitle,
            currentInfo,
            drawerAble,
            closeParent
        }
    },
}
</script>