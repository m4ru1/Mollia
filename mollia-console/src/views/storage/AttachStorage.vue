<template>
    <div class="attach-option">
        <PageContent class="page-content">
            <template #page-title>
                <h2> 图床管理 </h2>
            </template>
            <template #page-view>

            </template>
        </PageContent>
    </div>
</template>

<script setup>
import PageContent from '../../components/layout/PageContent.vue';
</script>

<style scoped>
.attach-option{
    display: flex;
    overflow: hidden;
    position: relative;
    background-color: var(--theme-pagecontent-bg-color);
}
</style>