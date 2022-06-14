<template>
    <el-drawer
      v-model="drawerAble"
      :title="currentTitle"
      @close="closeParent"
      >
        <span>Hi, there!</span>
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
        }
    },
    setup(props,context) {
        let drawerAble = ref(props.drawer);
        let currentTitle = ref(props.title);

        const closeParent = ()=>{
            context.emit("close");
        }
        
        watch(
            ()=>[props.drawer,props.title],
            (value) => {
                currentTitle.value = value[1];
                drawerAble.value = value[0];
            }
        );

        return {
            currentTitle,
            drawerAble,
            closeParent
        }
    },
}
</script>