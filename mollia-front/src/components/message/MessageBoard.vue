<template>
    <div class="message-board">
        <div class="messageBoard-header">
            <p>留言板</p>
        </div>
        <div class="messages">
            <MessageBox 
                v-for="(item, index) in testData"
                :id="index"
                :name="item.name"
                :avatar-u-r-l="item.avatarURL"
                :content="item.content"
                :date="item.date"
            ></MessageBox>
        </div>
    </div>
</template>

<script setup>
import MessageBox from './MessageBox.vue';

const testData = [{
    name: '李白',
    avatarURL: 'https://profile-avatar.csdnimg.cn/1c61abbaa7e5406c81afac8cbb68c676_qq_38236620.jpg!1',
    content: '你这网站建的不错',
    date: (new Date()).toLocaleDateString('zh-CN', { year: 'numeric', month: '2-digit', day: '2-digit' })
},
{
    name: '李白',
    avatarURL: 'https://profile-avatar.csdnimg.cn/1c61abbaa7e5406c81afac8cbb68c676_qq_38236620.jpg!1',
    content: '你这网站建的不错',
    date: (new Date()).toLocaleDateString('zh-CN', { year: 'numeric', month: '2-digit', day: '2-digit' })
},
{
    name: '李白',
    avatarURL: 'https://profile-avatar.csdnimg.cn/1c61abbaa7e5406c81afac8cbb68c676_qq_38236620.jpg!1',
    content: '你这网站建的不错',
    date: (new Date()).toLocaleDateString('zh-CN', { year: 'numeric', month: '2-digit', day: '2-digit' })
},
{
    name: '李白',
    avatarURL: 'https://profile-avatar.csdnimg.cn/1c61abbaa7e5406c81afac8cbb68c676_qq_38236620.jpg!1',
    content: '你这网站建的不错',
    date: (new Date()).toLocaleDateString('zh-CN', { year: 'numeric', month: '2-digit', day: '2-digit' })
},
{
    name: '李白',
    avatarURL: 'https://profile-avatar.csdnimg.cn/1c61abbaa7e5406c81afac8cbb68c676_qq_38236620.jpg!1',
    content: '你这网站建的不错',
    date: (new Date()).toLocaleDateString('zh-CN', { year: 'numeric', month: '2-digit', day: '2-digit' })
},
{
    name: '李白',
    avatarURL: 'https://profile-avatar.csdnimg.cn/1c61abbaa7e5406c81afac8cbb68c676_qq_38236620.jpg!1',
    content: '你这网站建的不错',
    date: (new Date()).toLocaleDateString('zh-CN', { year: 'numeric', month: '2-digit', day: '2-digit' })
},
{
    name: '李白',
    avatarURL: 'https://profile-avatar.csdnimg.cn/1c61abbaa7e5406c81afac8cbb68c676_qq_38236620.jpg!1',
    content: '你这网站建的不错',
    date: (new Date()).toLocaleDateString('zh-CN', { year: 'numeric', month: '2-digit', day: '2-digit' })
}]
</script>

<style>
.message-board {
    display: flex;
    flex-direction: column;
    position: relative;
    flex-shrink: 0;
    width: 100%;
    background-color: #111827;
    opacity: .8;
    border-radius: 20px;
    overflow: hidden;
}


.message-board > .messages{
    overflow: auto;/*超出部分自动显示滚动条，可以进行滚动*/
}

/*留言板滚动条美化，非标准语法，Firefox PC/Mobile以及Safari均不支持*/
.messages::-webkit-scrollbar{ /*滚动条宽度*/
    width: 2px;
    height: 1px;
}

.messages::-webkit-scrollbar-track{ /*滚动条轨道*/
    background-color: transparent;
    border-radius: 2em;
}

.messages::-webkit-scrollbar-thumb{ /*滚动条滑块*/
    background-color: rgb(255, 255, 255, .7);
    border-radius: 2em;
}

.messageBoard-header {
    position: sticky;
    display: flex;
    justify-content: center;
    top: 0; /*留言板标头固定*/
    margin-bottom: 10px;
    background-color: #111827;
    opacity: .97;
    border-radius: 10px;
    box-shadow: 0 0 1.1px white;
    padding: 32px 24px 24px 24px;
}

.messageBoard-header > p {
    margin: 0;
    font-size: 24px;
    line-height: 32px;
    font-weight: 700;
    opacity: .95;
    color: white;
    user-select: none;
}
</style>