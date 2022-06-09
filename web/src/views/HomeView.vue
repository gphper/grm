<template>
  <el-container>
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
export default {
  name: "HomeView",
  components: {
      MainHeader,
      MainIndex
  },
  setup(){
    document.title = "go redis manager";
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