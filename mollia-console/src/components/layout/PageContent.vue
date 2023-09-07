<template>
    <div class="page-content">
        <div class="page-title">
            <slot name="page-title"> </slot>
        </div>
        <div class="page-view">
            <slot name="page-view"> </slot>
        </div>
        <div class="page-footer">
            <PageFooter></PageFooter>
        </div>
    </div>
</template>

<script setup>
import PageFooter from './PageFooter.vue';

</script>

<style scoped>
.page-content{
    flex: 1;
    display: flex;
    flex-direction: column;
    position: relative;
    padding: 15px;
    overflow: auto;
    .page-title{
        color: var(--theme-pagetitle-text-color);
        padding: 0 15px 0 15px;
        user-select: none;
    }
    .page-view{
        display: flex;
        flex-direction: column;
        flex: 1;
        border-radius: 15px;
        background-color: var(--theme-pageview-bg-color);
    }
    .page-footer{
        padding-top: 10px;
    }
    &::-webkit-scrollbar{ /*滚动条宽度*/
        width: 4px;
        height: 1px;
    }

    &::-webkit-scrollbar-track{ /*滚动条轨道*/
        background-color: #1b1b1b;
    }

    &::-webkit-scrollbar-thumb{ /*滚动条滑块*/
        background-color: #10b981;
        border-radius: 90px;
    }
}
</style>