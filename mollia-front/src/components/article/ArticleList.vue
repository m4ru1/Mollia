<template>
    <div class="article-list">
        <ArticleEntry
            class="article-node"
            v-for="(entry, index) in testData"
            :key="index"
            :avatar-u-r-l="entry.avatar"
            :title="entry.title"
            :content="entry.content"
            :tag="entry.tag"
            :date="entry.date"
            :word-count="entry.wordCount"
        >
        </ArticleEntry>
    </div>
</template>

<script setup>
import ArticleEntry from './ArticleEntry.vue';
import {ref} from 'vue';

let testData = ref([
    {
        title: 'Some list item',
        content: "In the realm of dreams, where stars do shine, A world of wonder, where fantasies entwine. Where moonlight dances on the ocean's crest, And whispers of love fill the air with zest.",
        date: (new Date()).toLocaleDateString('zh-CN', { year: 'numeric', month: '2-digit', day: '2-digit' }),
        tag: '生活',
        avatar: 'https://profile-avatar.csdnimg.cn/1c61abbaa7e5406c81afac8cbb68c676_qq_38236620.jpg!1',
        wordCount: 125
    },
    {
        title: 'Some list item',
        content: "In the realm of dreams, where stars do shine, A world of wonder, where fantasies entwine. Where moonlight dances on the ocean's crest, And whispers of love fill the air with zest.",
        date: (new Date()).toLocaleDateString('zh-CN', { year: 'numeric', month: '2-digit', day: '2-digit' }),
        tag: '生活',
        avatar: 'https://profile-avatar.csdnimg.cn/1c61abbaa7e5406c81afac8cbb68c676_qq_38236620.jpg!1',
        wordCount: 125
    },
    {
        title: 'Some list item',
        content: "In the realm of dreams, where stars do shine, A world of wonder, where fantasies entwine. Where moonlight dances on the ocean's crest, And whispers of love fill the air with zest.",
        date: (new Date()).toLocaleDateString('zh-CN', { year: 'numeric', month: '2-digit', day: '2-digit' }),
        tag: '生活',
        avatar: 'https://profile-avatar.csdnimg.cn/1c61abbaa7e5406c81afac8cbb68c676_qq_38236620.jpg!1',
        wordCount: 125
    },
    {
        title: 'Some list item',
        content: "In the realm of dreams, where stars do shine, A world of wonder, where fantasies entwine. Where moonlight dances on the ocean's crest, And whispers of love fill the air with zest.",
        date: (new Date()).toLocaleDateString('zh-CN', { year: 'numeric', month: '2-digit', day: '2-digit' }),
        tag: '生活',
        avatar: 'https://profile-avatar.csdnimg.cn/1c61abbaa7e5406c81afac8cbb68c676_qq_38236620.jpg!1',
        wordCount: 125
    },
    {
        title: 'Some list item',
        content: "In the realm of dreams, where stars do shine, A world of wonder, where fantasies entwine. Where moonlight dances on the ocean's crest, And whispers of love fill the air with zest.",
        date: (new Date()).toLocaleDateString('zh-CN', { year: 'numeric', month: '2-digit', day: '2-digit' }),
        tag: '生活',
        avatar: 'https://profile-avatar.csdnimg.cn/1c61abbaa7e5406c81afac8cbb68c676_qq_38236620.jpg!1',
        wordCount: 125
    },
    {
        title: 'Some list item',
        content: "In the realm of dreams, where stars do shine, A world of wonder, where fantasies entwine. Where moonlight dances on the ocean's crest, And whispers of love fill the air with zest.",
        date: (new Date()).toLocaleDateString('zh-CN', { year: 'numeric', month: '2-digit', day: '2-digit' }),
        tag: '生活',
        avatar: 'https://profile-avatar.csdnimg.cn/1c61abbaa7e5406c81afac8cbb68c676_qq_38236620.jpg!1',
        wordCount: 125
    },
    {
        title: 'Some list item',
        content: "In the realm of dreams, where stars do shine, A world of wonder, where fantasies entwine. Where moonlight dances on the ocean's crest, And whispers of love fill the air with zest.",
        date: (new Date()).toLocaleDateString('zh-CN', { year: 'numeric', month: '2-digit', day: '2-digit' }),
        tag: '生活',
        avatar: 'https://profile-avatar.csdnimg.cn/1c61abbaa7e5406c81afac8cbb68c676_qq_38236620.jpg!1',
        wordCount: 125
    },
    {
        title: 'Some list item',
        content: "In the realm of dreams, where stars do shine, A world of wonder, where fantasies entwine. Where moonlight dances on the ocean's crest, And whispers of love fill the air with zest.",
        date: (new Date()).toLocaleDateString('zh-CN', { year: 'numeric', month: '2-digit', day: '2-digit' }),
        tag: '生活',
        avatar: 'https://profile-avatar.csdnimg.cn/1c61abbaa7e5406c81afac8cbb68c676_qq_38236620.jpg!1',
        wordCount: 125
    },
    {
        title: 'Some list item',
        content: "In the realm of dreams, where stars do shine, A world of wonder, where fantasies entwine. Where moonlight dances on the ocean's crest, And whispers of love fill the air with zest.",
        date: (new Date()).toLocaleDateString('zh-CN', { year: 'numeric', month: '2-digit', day: '2-digit' }),
        tag: '生活',
        avatar: 'https://profile-avatar.csdnimg.cn/1c61abbaa7e5406c81afac8cbb68c676_qq_38236620.jpg!1',
        wordCount: 125
    },
    {
        title: 'Some list item',
        content: "In the realm of dreams, where stars do shine, A world of wonder, where fantasies entwine. Where moonlight dances on the ocean's crest, And whispers of love fill the air with zest.",
        date: (new Date()).toString(),
        tag: '生活',
        avatar: 'https://profile-avatar.csdnimg.cn/1c61abbaa7e5406c81afac8cbb68c676_qq_38236620.jpg!1',
        wordCount: 125
    },
    {
        title: 'Some list item',
        content: "In the realm of dreams, where stars do shine, A world of wonder, where fantasies entwine. Where moonlight dances on the ocean's crest, And whispers of love fill the air with zest.",
        date: (new Date()).toString(),
        tag: '生活',
        avatar: 'https://profile-avatar.csdnimg.cn/1c61abbaa7e5406c81afac8cbb68c676_qq_38236620.jpg!1',
        wordCount: 125
    },
    {
        title: 'Some list item',
        content: "In the realm of dreams, where stars do shine, A world of wonder, where fantasies entwine. Where moonlight dances on the ocean's crest, And whispers of love fill the air with zest.",
        date: (new Date()).toString(),
        tag: '生活',
        avatar: 'https://profile-avatar.csdnimg.cn/1c61abbaa7e5406c81afac8cbb68c676_qq_38236620.jpg!1',
        wordCount: 125
    },
])
</script>

<style>

.article-node{
    list-style-type: none;
    margin: 0 0 20px 0;
}

</style>