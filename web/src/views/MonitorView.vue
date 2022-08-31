<template>
    <div class="common-layout">
    <el-container>
      <el-header>
        <el-row>
            <el-col :span="8">
                <el-switch
                    v-model="showswitch"
                    class="mb-2"
                    style="--el-switch-on-color: #13ce66; --el-switch-off-color: #ff4949"
                    active-text="表格"
                    inactive-text="图表"
                />
            </el-col>

            <el-col :span="4" :offset="6">
                <el-form-item label="是否自动更新">
                    <el-switch
                        @change="autoChange"
                        v-model="autoupdate"
                        inline-prompt
                        active-text="是"
                        inactive-text="否"
                    />
                </el-form-item>
            </el-col>

            <el-col :span="3">
                <el-form-item label="更新间隔">
                    <el-select @change="intervalChange" v-model="interval" placeholder="间隔">
                        <el-option label="5s" value="5" />
                        <el-option label="10s" value="10" />
                        <el-option label="30s" value="30" />
                        <el-option label="1min" value="60" />
                        <el-option label="5min" value="300" />
                    </el-select>
                </el-form-item>
            </el-col>

            <el-col :span="2" :offset="1">
                <el-button type="success" title="重载" circle><i class="iconfont icon-zhongqi"></i></el-button>
            </el-col>

        </el-row>
      </el-header>
      <el-main>
            <div v-show="!showswitch">chart</div>
            <div v-show="showswitch">data list</div>
      </el-main>
    </el-container>
  </div>
</template>

<script>
import { onMounted, ref } from '@vue/runtime-core'
import { useRoute } from 'vue-router'
import { serInfo } from "@/api/index.js"
export default {
    name:"MonitorView",
    setup() {
        const route = new useRoute()
        let timer = null

        const intervalChange = (val)=>{
            clearTimeout(timer)
            timer = setInterval(()=>{
                console.log(new Date().getMinutes(),new Date().getSeconds())
            },val * 1000)
        }

        const autoChange = (val) => {
            console.log(val)
            if(val == false){
                clearTimeout(timer)
                timer = null
            }else{
                timer = setInterval(()=>{
                    console.log(new Date().getMinutes(),new Date().getSeconds())
                },interval.value * 1000)
            }
        }

       let showswitch =  ref(false)
       let autoupdate = ref(true)
       let interval = ref(5)


       if(autoupdate.value){
            timer = setInterval(()=>{
                console.log(new Date().getMinutes(),new Date().getSeconds())
            },interval.value * 1000)
       }

       onMounted(() => {
            // 打印
            console.log(route.params.sk)
            serInfo({key:"8e7eca1f7079c6a7861b3dcaa2dc8c05"}).then((res)=>{
                console.log(res)
            })
       }) 

       return {
            showswitch,
            interval,
            autoupdate,
            autoChange,
            intervalChange
       }
    },
}
</script>

<style scoped>
.el-header{
    background-color:  #EBEDF0;
}
</style>