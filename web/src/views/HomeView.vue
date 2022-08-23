<template>
  <el-container v-loading="loading" element-loading-background="rgba(0, 0, 0, 0)">
      <el-header>
        <MainHeader></MainHeader>
      </el-header>
      <el-main>
        <MainIndex></MainIndex>
      </el-main>
  </el-container>
	
</template>

<script>
import MainHeader from "@/components/MainHeader.vue";
import MainIndex from "@/components/MainIndex.vue";
import { ref } from '@vue/reactivity';
import { useStore } from "vuex";
import { onMounted, watch } from '@vue/runtime-core';
export default {
  name: "HomeView",
  components: {
      MainHeader,
      MainIndex
  },
  setup(){
    let loading = ref(false);
    const store = useStore();
    
     watch(
          () => store.state.loadReq,
          (value) => {
              loading.value = value
          }
      )

    onMounted(()=>{
      document.title = "go redis manager";
    })  

    return {
      loading
    }
  }
};
</script>

<style type="text/css" scoped>
    .el-header{
    /* background-color: #EBEEF5; */
      line-height: 50px;
      height: 55px;
      border-bottom: 1px solid  #a0cfff;
    }

    .el-main{
      padding:0px;
    }
</style>