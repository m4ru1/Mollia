<template>
    <div class="message-box">
        <div class="message-avatar">
            <img :src="props.avatarURL">
        </div>
        <div class="message-content">
            <div class="message-header">
                <div class="name">{{ props.name }}</div>
                <div class="like-checkbox">
                    <input :id="`star-${props.id}`" type="checkbox" style="display: none;" v-model="hasLiked">
                    <label :for="`star-${props.id}`">
                        <!--可以控制svg的fill字段，#1296db即为我们想要的蓝色，#333333就是默认灰色-->
                        <svg v-show="true" t="1691653090963" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="4048" width="200" height="200"><path d="M840.88 330H631.67c-11.17 0-19.3-10.6-16.41-21.38l23.16-86.43c14.45-53.93-17.55-109.36-71.48-123.81A98.616 98.616 0 0 0 541.36 95c-43.88 0-84.3 29.38-96.17 73.66l-19.8 73.88C411.57 294.12 364.82 330 311.41 330H150c-17.67 0-32 14.33-32 32v448c0 17.67 14.33 32 32 32h618.88c30.8 0 57.23-21.94 62.9-52.21l72-384c7.39-39.38-22.82-75.79-62.9-75.79zM182 394h93.63L248.2 778H182V394z m586.88 384H312.37l27.6-386.39c69.97-11.15 128.47-62.44 147.25-132.51l19.8-73.88c4.07-15.2 18.52-26.22 34.35-26.22 3.01 0 6.04 0.4 9.01 1.2 19.76 5.29 31.52 25.68 26.23 45.43l-23.16 86.43c-6.57 24.52-1.47 50.13 13.98 70.26 15.45 20.14 38.87 31.69 64.25 31.69h209.21L768.88 778z" :fill="hasLiked ? '#1296db' : '#aaaaaa'" p-id="4049"></path></svg>
                    </label>
                </div>
            </div>
            <div class="message-line">
                {{ props.content }}
            </div>
            <div class="message-line time">
                {{ props.date }}
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref } from 'vue';

const props = defineProps({
    id: String | Number,
    name: String,
    avatarURL: String,
    content: String,
    date: String,
    liked: Boolean
})
const hasLiked = ref(props.liked);
</script>

<style>
.message-box{
    display: flex;
    padding: 10px 10px 10px 10px;
    border-bottom: 1px solid #444;
    background-color: #111827;
    transition: all .2s ease;
}

.message-box:hover{
    background-color: #131a29;
}

.message-avatar{
    flex-grow: 0;
}

.message-avatar > img{
    width: 35px;
    height: 35px;
    border-radius: 100%;
    user-select: none;
    -webkit-user-drag: none;
}

.message-content{
    flex-grow: 1;
    display: flex;
    margin: auto 10px auto 10px;
    flex-direction: column;
    justify-content: center;
}

.message-header{
    display: flex;
}

.like-checkbox, .name{
    display: inline;
}

.name{
    color: white;
    font-size: 16px;
    font-weight: 600;
    flex-grow: 1;
    user-select: none;
}

svg{
    width: 20px;
    height: 20px;
    text-align: right;
}

.message-line{
    color: white;
    -webkit-user-drag: none;
}


.like-checkbox{
    text-align: right;
}

.time{
    text-align: right;
    user-select: none;
}
</style>