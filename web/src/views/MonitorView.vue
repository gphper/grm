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
            <div v-show="!showswitch">
                <el-row>
                    <el-col :span="8" style="height:350px;">
                        <div style="width:100%;height:100%" id="chart1"></div>
                    </el-col>
                    <el-col :span="8" style="height:350px;">
                        <div style="width:100%;height:100%" id="chart2"></div>
                    </el-col>
                    <el-col :span="8" style="height:350px;">
                        <div style="width:100%;height:100%" id="chart3"></div>
                    </el-col>
                    <el-col :span="8" style="height:350px;">
                        <div style="width:100%;height:100%" id="chart4"></div>
                    </el-col>
                    <el-col :span="8" style="height:350px;">
                        <div style="width:100%;height:100%" id="chart5"></div>
                    </el-col>
                    <el-col :span="8" style="height:350px;">
                        <div style="width:100%;height:100%" id="chart6"></div>
                    </el-col>
                </el-row>
            
            </div>
            <div v-show="showswitch">data list</div>
      </el-main>
    </el-container>
  </div>
</template>

<script>
import { onMounted, ref } from '@vue/runtime-core'
import { useRoute } from 'vue-router'
import { serInfo } from "@/api/index.js"
import * as echarts from "echarts";
export default {
    name:"MonitorView",
    setup() {
        const route = new useRoute()
        let timer = null
        let xtime = []
        //内存
        let memuse = []
        let memChart = null
        let memChartOption = []
        const memChartChange = (used_memory)=>{
            if(memuse.length < 100){
                memuse.push(used_memory)
            }else{
                memuse.shift()
                memuse.push(used_memory)
            }
            memChartOption.xAxis.data = xtime
            memChartOption.series[0].data = memuse
            memChart.setOption(memChartOption)
        }


        //进口流量
        let input_kbps = []
        let inputChart = null
        let inputChartOption = []
        const inputChartChange = (input_kbps_val)=>{
            if(input_kbps.length < 100){
                input_kbps.push(input_kbps_val)
            }else{
                input_kbps.shift()
                input_kbps.push(input_kbps_val)
            }
            inputChartOption.xAxis.data = xtime
            inputChartOption.series[0].data = input_kbps
            inputChart.setOption(inputChartOption)
        }

        //出口流量
        let output_kbps = []
        let outputChart = null
        let outputChartOption = []
        const outputChartChange = (output_kbps_val)=>{
            if(output_kbps.length < 100){
                output_kbps.push(output_kbps_val)
            }else{
                output_kbps.shift()
                output_kbps.push(output_kbps_val)
            }
            outputChartOption.xAxis.data = xtime
            outputChartOption.series[0].data = output_kbps
            outputChart.setOption(outputChartOption)
        }

        const trans = (size) => {
            let kb = 1024
            let unit = ['B','KB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB']
            let i = Math.floor(Math.log(size) / Math.log(kb))
            if(size > 0){
                return (size / Math.pow(kb, i)).toPrecision(3) + ' ' + unit[i];
            }else{
                return '0B'
            } 
        }

        

        const getInfo = ()=>{
            serInfo({key:"8e7eca1f7079c6a7861b3dcaa2dc8c05"}).then((res)=>{
                if(xtime.length < 100){
                    xtime.push(res.data.time)
                }else{
                    xtime.shift()
                    xtime.push(res.data.time)
                }
                
                //更新内存
                memChartChange(res.data.info.Memory.used_memory)
                //更新进口流量
                inputChartChange(res.data.info.Stats.instantaneous_input_kbps)
                //更新出口流量
                outputChartChange(res.data.info.Stats.instantaneous_output_kbps)
            })
        }

        const intervalChange = (val)=>{
            clearTimeout(timer)
            timer = setInterval(()=>{
                getInfo()
            },val * 1000)
        }

        const autoChange = (val) => {
            if(val == false){
                clearTimeout(timer)
                timer = null
            }else{
                timer = setInterval(()=>{
                    getInfo()
                },interval.value * 1000)
            }
        }

       let showswitch =  ref(false)
       let autoupdate = ref(true)
       let interval = ref(5)


       if(autoupdate.value){
            timer = setInterval(()=>{
                getInfo()
            },interval.value * 1000)
       }

       onMounted(() => {
            // 打印
            console.log(route.params.sk)
            xtime = []
            
            memChart = echarts.init(document.getElementById('chart1'));
            memChartOption = {
                title: {
                    text: 'Redis内存使用'
                },
                tooltip: {
                    trigger: 'axis',
                    formatter:(params)=>{
                        return params[0].axisValueLabel+":"+ trans(params[0].data);
                    },
                },
                xAxis: {
                    type: 'category',
                    data: []
                },
                yAxis: {
                    type: 'value',
                    axisLabel:{
                        formatter:(value) => {
                            return trans(value)
                        }
                    }
                },
                series: [
                    {
                        data: [],
                        type: 'line',
                        smooth: true
                    }
                ]
            };
            memChart.setOption(memChartOption);


            inputChart = echarts.init(document.getElementById('chart2'));
            inputChartOption = {
                title: {
                    text: 'Redis进口流量'
                },
                tooltip: {
                    trigger: 'axis',
                    formatter:(params)=>{
                        return params[0].axisValueLabel+":"+ params[0].data + "kbps";
                    },
                },
                xAxis: {
                    type: 'category',
                    data: []
                },
                yAxis: {
                    type: 'value',
                    axisLabel:{
                        formatter:(value) => {
                            return value + "kbps"
                        }
                    }
                },
                series: [
                    {
                        data: [],
                        type: 'line',
                        smooth: true
                    }
                ]
            };
            inputChart.setOption(inputChartOption);

            outputChart = echarts.init(document.getElementById('chart3'));
            outputChartOption = {
                title: {
                    text: 'Redis出口流量'
                },
                tooltip: {
                    trigger: 'axis',
                    formatter:(params)=>{
                        return params[0].axisValueLabel+":"+ params[0].data + "kbps";
                    },
                },
                xAxis: {
                    type: 'category',
                    data: []
                },
                yAxis: {
                    type: 'value',
                    axisLabel:{
                        formatter:(value) => {
                            return value + "kbps"
                        }
                    }
                },
                series: [
                    {
                        data: [],
                        type: 'line',
                        smooth: true
                    }
                ]
            };
            outputChart.setOption(outputChartOption);
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