<template>
    <div>
        <JsonViewer :value="currentData" boxed sort />
    </div>
</template>

<script>
import {JsonViewer} from "vue3-json-viewer"
import "vue3-json-viewer/dist/index.css";
import { computed } from '@vue/runtime-core';

export default {
    name:"OneJson",
    props:{
        data:{
            type:String
        }
    },
    components:{
        JsonViewer
    },
    setup(props) {

        let currentData = computed({
            get:()=>{
                try {
                    let tmp = JSON.parse(props.data);
                    return tmp;
                } catch(error) {
                    return props.data;
                }
            },
        });

        return {
            currentData
        }
    },
}
</script>