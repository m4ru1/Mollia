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
    display: flex;
    flex-direction: column;
    overflow: auto;
    position: relative;
    padding: 15px;
    .page-title{
        color: var(--theme-pagetitle-text-color);
        padding: 0 15px 0 15px;
        user-select: none;
    }
    .page-view{
        flex: 1;
        border-radius: 15px;
        background-color: var(--theme-pageview-bg-color);
    }
    .page-footer{
        padding-top: 10px;
    }
}
</style>