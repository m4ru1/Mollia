<template>
<div class="pagination-board">
    <span :class="['shift-button', {'shift-disabled': isPrevDisabled}]" @click="lastPage">
        <svg t="1693212457629" class="shift-button-icon" viewBox="0 0 1198 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="11058" width="200" height="200">+
            <g>
                <path d="M1086.53-5.345l-463.494 515.362 463.494 514.257-81.663 73.936-529.707-588.196 529.706-589.298 81.663 73.936z" p-id="11059"></path>
                <path d="M691.456-5.345l-462.391 515.362 462.391 514.257-81.663 73.936-529.707-588.196 529.706-589.298 81.663 73.936z" p-id="11060"></path>
            </g> 
        </svg>
    </span>
    <ul class="page-box">
        <li :class="['page-item', item.type, {'curr-page': currPage === item.content}]" 
            v-for="(item, index) in paginationList" :key="index"
            @click="item.type === 'number'? changePage(item.content) : changeMenu(item.type)">
            {{ item.content }}
        </li>
    </ul>
    <span :class="['shift-button', {'shift-disabled': isNextDisabled}]" @click="nextPage">
        <svg t="1693212494871" class="shift-button-icon" viewBox="0 0 1198 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="1400" width="200" height="200">
            <g> <!-- g元素可以使SVG中的每一个PATH可以同时应用样式 -->
                <path d="M97.46 1024.012l463.494-515.362-463.494-514.257 81.663-73.936 529.707 588.196-529.706 589.298-81.663-73.936z" p-id="1401"></path>
                <path d="M492.534 1024.012l462.391-515.362-462.391-514.257 81.663-73.936 529.707 588.196-529.706 589.298-81.663-73.936z" p-id="1402"></path>
            </g>
        </svg>
    </span>
</div>
</template>

<script setup>
// import { font-awesome-icon } from 'font-awesome';
import { ref, defineProps, onMounted } from 'vue';

const props = defineProps({
    totalNumber: {  // 总条目数量
        type: Number,
        required: true
    },
    perPage: {      // 单页条目数量
        type: Number,
        default: 10
    },
    pagerCount: {    // 分页器显示的条目数量
        type: Number,
        default: 7
    }
});

let paginationList = ref([]);                                   // 分页对象列表
let isPrevDisabled, isNextDisabled = true;                      // 页面切换按钮是否禁用
let totalPage = Math.ceil(props.totalNumber / props.perPage);   // 页数
let currPage = ref('1');                                        // 当前页码
let pagerCount = props.pagerCount;                              // 分页器显示的条目数量                             


function leftToEnd(){   // 初始首页
    if(totalPage > pagerCount){
        let currLen = paginationList.value.length;
        for(let i = 0; i < currLen; i++){
            paginationList.value.pop();
        }
        for(let i = 1; i < pagerCount; i++){
            paginationList.value.push({type: 'number', content: `${i}`});
        }
        paginationList.value.push({type: 'rightBtn', content: '...'});
        paginationList.value.push({type: 'number', content: `${totalPage}`});
    }else{
        for(let i = 1; i <= totalPage; i++){
            paginationList.value.push({type: 'number', content: `${i}`});
        }
    }
}

function rightToEnd(){  // 初始末页
    if(totalPage > pagerCount){
        let currLen = paginationList.value.length;
        for(let i = 0; i < currLen; i++){
            paginationList.value.pop();
        }
        paginationList.value.push({type: 'number', content: '1'});
        paginationList.value.push({type:'leftBtn', content: '...'});
        for (let i = totalPage - (pagerCount - 1) + 1; i <=  totalPage; i++){
            paginationList.value.push({type: 'number', content: `${i}`});
        }
    }else{
        for(let i = 1; i <= totalPage; i++){
            paginationList.value.push({type: 'number', content: `${i}`});
        }
    }
}

function changeMenu(type){
    if(type === 'leftBtn'){
        leftChangeMenu();
    }else{
        rightChangeMenu();
    }
}

// 1. 分页初始化
leftToEnd();

// 2. 左切换菜单
function leftChangeMenu(allowChangePage = true){
    let currLen = paginationList.value.length;
    if (currLen > pagerCount){
        let midStart = parseInt(paginationList.value[2].content); // 中间动态菜单的首个页码
        if(currLen === pagerCount + 1 && paginationList.value[1].type === 'leftBtn'){ // 仅有左切换按钮
            if(midStart - 2 >= 4){
                for(let i = 0; i < currLen - 2; i++){
                    paginationList.value.pop();
                }
                for(let i = midStart - 2; i <= (midStart - 2) + (pagerCount - 2) - 1; i++){
                    paginationList.value.push({type: 'number', content: `${i}`});
                }
            }else{
                for(let i = 0; i < currLen; i++){
                    paginationList.value.pop();
                }
                for(let i = 1; i <= pagerCount - 1; i++){
                    paginationList.value.push({type: 'number', content: `${i}`});
                }
            }
            paginationList.value.push({type: 'rightBtn', content: '...'});
            paginationList.value.push({type: 'number', content: `${totalPage}`});
        }else if(currLen === pagerCount + 2){ // 同时存在左切换与右切换按钮
            if(midStart >= pagerCount + 2){ // 只改变中间的浮动菜单即可
                for(let i = 0; i < pagerCount - 2; i++){
                    paginationList.value[i + 2].content = `${(midStart + i) - (pagerCount - 2)}`;
                }
            }else{  // 左侧未展示页码不支撑再次切换，因此
                for(let i = 0; i < currLen; i++){
                    paginationList.value.pop();
                }
                for(let i = 1; i <= pagerCount - 1; i++){
                    paginationList.value.push({type: 'number', content: `${i}`});
                }
                paginationList.value.push({type: 'rightBtn', content: '...'});
                paginationList.value.push({type: 'number', content: `${totalPage}`});
            }
        }
        if(allowChangePage){
            changePage(paginationList.value[Math.floor(paginationList.value.length / 2)].content);
        }  
    }
}

// 3. 右切换菜单
function rightChangeMenu(allowChangePage = true){
    let currLen = paginationList.value.length;
    if(currLen > pagerCount){
        let midEnd = parseInt(paginationList.value[currLen - 3].content);
        if(currLen === pagerCount + 1 && paginationList.value[currLen - 2].type === 'rightBtn'){ // 仅有右切换按钮
            for(let i = 0; i < currLen - 1; i++){
                paginationList.value.pop();
            }
            paginationList.value.push({type: 'leftBtn', content: '...'});
            if(midEnd + 2 <= totalPage - 3){
                for(let i = (midEnd + 2) - (pagerCount - 2) + 1; i <= midEnd + 2; i++){
                    paginationList.value.push({type: 'number', content: `${i}`});
                }
                paginationList.value.push({type: 'rightBtn', content: '...'});
                paginationList.value.push({type: 'number', content: `${totalPage}`});
            }else{
                for(let i = totalPage - (pagerCount - 1) + 1; i <= totalPage; i++){
                    paginationList.value.push({type: 'number', content: `${i}`});
                }
            }
        }else if(currLen === pagerCount + 2){ // 同时存在左切换与右切换按钮
            if(midEnd + (pagerCount - 2) < totalPage - 3){
                for(let i = 0; i < pagerCount - 2; i++){
                    paginationList.value[i + 2] = {
                        type: 'number',
                        content: `${parseInt(paginationList.value[i + 2].content) + (pagerCount - 2)}`
                    };
                }
            }else{
                for(let i = 0; i < currLen - 2; i++){
                    paginationList.value.pop();
                }
                for(let i = totalPage - (pagerCount - 2); i <= totalPage; i++){
                    paginationList.value.push({type: 'number', content: `${i}`});   
                }
            }
        }
        if(allowChangePage){
            changePage(paginationList.value[Math.floor(paginationList.value.length / 2)].content);
        }
    }
}

// 4. 前一页
function lastPage(){
    if (parseInt(currPage.value) > 1){
        currPage.value = `${parseInt(currPage.value) - 1}`;
    }
    // 如果currPage不在当前菜单中，则自动向左切换菜单
    let isExist = false;
    for(let i = 0; i < paginationList.value.length; i++){
        if(paginationList.value[i].content === currPage.value){
            isExist = true;
            break; 
        }
    }
    if (!isExist){
        leftChangeMenu(false);
    }
}

// 5. 后一页
function nextPage(){
    if (parseInt(currPage.value) < totalPage){
        currPage.value = `${parseInt(currPage.value) + 1}`;
    }
    // 如果currPage不在当前菜单中，则自动向右切换菜单
    let isExist = false;
    for(let i = 0; i < paginationList.value.length; i++){
        if(paginationList.value[i].content === currPage.value){
            isExist = true;
            break; 
        }        
    }
    if (!isExist){
        rightChangeMenu(false);
    }
}

// 6. 触发换页
function changePage(targetPage){
    if (parseInt(targetPage) && (1 <= parseInt(targetPage) <= paginationList.value.length)){
        currPage.value = targetPage;
    }
    if (currPage.value == '1'){
        // 对于选中首页为标签的，向左切换到没有该按钮为止（直接一点就是重置菜单）
        leftToEnd();
    }else if(currPage.value === `${totalPage}`){
        // 同上
        rightToEnd();
    }
}
</script>

<style>
.pagination-board{
    width: 100%;
    display: flex;
    flex-direction: row;
    justify-content: center;
    font-size: 14px;
    border-radius: 10px;
    background-color: rgba(255, 255, 255, .6);
    -webkit-user-drag: none;
    .shift-button{
        height: 32px;
        width: 32px;
        margin: auto 0 auto 0;
        border-radius: 5px;
        display: flex;
        justify-content: center;
        align-items: center;
        user-select: none;
        cursor: pointer;
        &, & .shift-button-icon{
            transition: all .1s ease-in-out;
        }
        .shift-button-icon{
            width: 60%;
            height: 60%;
            fill: black;
        }
        &:hover .shift-button-icon{
            width: 300;
            fill: #10b981;
        }
        &:active{
            background-color: rgba(0, 0, 0, .3);
        }
    }
    .page-box{
        display: flex;
        flex-direction: row;
        align-items: center;
        justify-content: center;
        height: 36px;
        padding: 0px 20px;
        .page-item{
            width: 24px;
            height: 32px;
            display: flex;
            justify-content: center;
            align-items: center;
            padding: 0px 4px;
            margin: 0 10px;
            border-radius: 5px;
            list-style: none;
            user-select: none;
            background-color: rgba(220, 220, 220, .5);
            color: black;
            font-weight: 800;
            cursor: pointer;
            transition: all .1s ease-in;
            &:hover{
                color: #10b981;
                font-weight: 800;
            }
            &:active{
                background-color: rgba(0, 0, 0, .3);
            }
        }
        .curr-page{
            height: 36px;   /* 当前页item将会比其他item的边长大2px，
                               这里用transition处理，将会有Q弹切换动效 */
            padding: 0px 6px;
            box-shadow: 0px 0px 5px #10b981;
        }
    }
}

@media screen and (max-width: 800px){
    .pagination-board{
        .shift-button{
            display: none;
        }
    }
}
</style>