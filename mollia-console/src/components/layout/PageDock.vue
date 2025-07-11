<template>
<div class="page-dock">
    <div class="dock-content">
        <PageTag v-for="(item, index) in pageStore.pages" :key="index"
                :page-name="pageStore.pageName(item)"
                @tag-close="startTimer">
        </PageTag>
        <PageTag v-for="item in [1,2,3,4,5,6,7,8]"
        :page-name="`测试文章${item}`"></PageTag>
    </div>
</div>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue';
import { usePageStore } from '../../../stores/PageStore';
import PageTag from './PageTag.vue';

let pageStore = usePageStore();
let translateX = ref(0);

// 归位计时器
let timer;
function cancleTimer(){
    // 取消定时
    if(timer){
        clearTimeout(timer);
    }
}
function startTimer(duration=200){
    // 计时
    timer = setTimeout(()=>{
        let app = document.querySelector('.app');
        let nav = document.querySelector('.nav-board');
        let dock = document.querySelector('.page-dock')
        let dockContent = document.querySelector('.page-dock .dock-content');
        if (app && nav && dock && dockContent){
            let rightWidth = app.clientWidth - nav.clientWidth;
            let dockWidth = dock.clientWidth;
            if((translateX.value + dockWidth) < rightWidth){
                translateX.value = rightWidth - dockWidth;
            }else{
                translateX.value = 0;
            }
            dockContent.style.transition = 'all .5s ease-out';
            dockContent.style.translate = `${translateX.value}px 0px`;
            dockContent.style.transition = 'all .1s ease';
        }
    }, duration);
}


onMounted(() => {
    // 滚轮事件检测，自适应调整页面dock栏的水平位移
    addEventListener('wheel', (event) => {
        let clientX = event.clientX;    // event.screenX的取值范围从左到右是(-width, 0)
        let clientY = event.clientY;
        let nav = document.querySelector('.nav-board');
        let appHeader = document.querySelector('.app-header');
        let highRow = document.querySelector('.high-row');

        if(clientX > nav.clientWidth   // 区域限制
            && clientY > highRow.clientHeight 
            && clientY < appHeader.clientHeight){

            let app = document.querySelector('.app');
            let dock = document.querySelector('.page-dock')
            let dockContent = document.querySelector('.page-dock .dock-content');
            if (app && nav && dock && dockContent){
                // 自适应仿弹簧Dock栏逻辑设计
                // 效果1: 离标定点越远，滚轮能够移动的幅度越小（俗称: 滚出界后，越滚越滚不动）
                //        求解: y = distance * (-(x / maxOffset) ^ 2)
                //        y: 本次移动距离, x: 当前偏离标定点距离，
                //        maxOffset: 最大偏离标定点距离, distance: 正常移动距离(即不加限制时)
                // 效果2: 滚轮一旦停止，将会Q弹回滚至标定点 
                //        思路：每一次滚动都会刷新计时器，当计时器为0时，
                //        检测translateX的变化，如果出界，则辅以transition特定的mode就能实现自适应Q弹回滚
                // 效果3：每一个tag在被close掉时触发tag-close事件，
                //        启动一个上述计时器任务用于更新dock栏的偏移位置

                // DelatY 为负值时（即向左滚动）进行自适应
                let maxOffset;
                if(event.deltaY < 0){
                    // 计算标定点
                    maxOffset = Math.abs((app.clientWidth - nav.clientWidth) - dock.clientWidth) + 50;
                    let distance = event.deltaY * 0.3;
                    var y;
                    if(translateX.value > 0){
                        y = -distance;
                    }else{
                        y = - Math.abs(distance) * Math.pow((translateX.value / maxOffset), 2) + Math.abs(distance);
                    }
                    translateX.value -= y;
                    dockContent.style.translate = `${translateX.value}px 0px`;
                    if(translateX.value < 0 && (translateX.value + dock.clientWidth) < (app.clientWidth - nav.clientWidth)){
                        cancleTimer();
                        startTimer();   // 超过右侧标定点，倒计时结束后进行归位，该计时效果可以被覆盖
                    }
                }
                // DeltaY 为正值时（即向右滚动）进行自适应
                else{
                    // 计算标定点
                    maxOffset = 50;
                    let distance = event.deltaY * 0.3;
                    if(translateX.value < 0){
                        y = distance;
                    }else{
                        y = - Math.abs(distance) * Math.pow((translateX.value / maxOffset), 2) + Math.abs(distance);
                    }
                    translateX.value += y;
                    dockContent.style.translate = `${translateX.value}px 0px`;
                    if(translateX.value > 0){
                        cancleTimer(); 
                        startTimer();   // 超过左侧标定点，倒计时结束后进行归位，该计时效果可以被覆盖
                    }
                }
            }
        }
    });
})

// 订阅当前页面变动，若当前页面的Tag不处于PageDock可视部分中，将PageDock自动滚动到位
watch(pageStore.currPage, (state) => {
    // TODO: 该动效逻辑后面再补坑
});

</script>

<style scoped>
.page-dock{
    position: relative;
    overflow: hidden;
    --el-text-color-regular: black;
    padding: 5px 30px;
    .dock-content{
        display: flex;
        overflow: hidden;
        transition: all .05s ease-in-out;
    }
}
</style>