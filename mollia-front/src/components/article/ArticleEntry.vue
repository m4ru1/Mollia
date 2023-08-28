<template>
    <div class="article-entry">
        <div class="avatar-area">
            <img :src="props.avatarURL">
        </div>
        <div class="main-area">
            <div class="title-area">
                <a href="h">{{ props.title }}</a>
            </div>
            <div class="info-area">
                <i class="date"> {{ props.date }} </i>
                <span class="divider">|</span>
                <i class="tag"> {{ props.tag }} </i>
                <span class="divider">|</span>
                <i class="word-count"> {{ props.wordCount.toString() + (props.wordCount > 1 ? ' words' : ' word') }} </i>
                <span class="divider">|</span>
                <i class="read-time"> {{ readTime.toString() + (readTime > 1 ? ' minitus' : ' minitu') }} </i>
            </div>
            <div class="content-area">
                {{ props.content }}
            </div>
        </div>
    </div>
</template>

<script setup>

// 根据下列信息渲染一个文章条目
const props = defineProps({
    avatarURL: String,
    title: String,
    content: String,
    tag: String,
    date: String,
    wordCount: Number
})

const readTime = Math.max(1, Math.floor(props.wordCount / 150));

</script>

<style>
.article-entry{
    display: flex;
    border-radius: 10px;
    justify-content: start;
    background-color: rgba(255, 255, 255, .7);
    transition: background-color .5s ease;
}

.article-entry:hover{
    background-color: rgba( 245, 245, 245, .9);
}

/*文章图像样式*/
.avatar-area{
    display: flex;
    border-radius: inherit;
    box-shadow: 3px 2px 2px #ccc; 
    flex-grow: 0;
    overflow: hidden;
}

.avatar-area > img{
    user-select: none;
    -webkit-user-drag: none;
    transition: transform .8s ease;
}

.avatar-area > img:hover{
    transform: scale(1.3, 1.3) translate(20px, -20px);
}

.main-area{
    display: flex;
    flex-grow: 5;
    padding: 0 20px 20px 20px;
    flex-direction: column;
    justify-content: start;
    text-align: center;
}

/*文章标题样式*/
.title-area{
    display: flex;
    justify-content: center;
    margin-top: 30px;
    overflow: hidden;
}

.title-area > a{
    position: relative;
    display: block;
    font-size: 35px;
    max-width: 710px;
    max-height: 80px;
    color: black;
    font-size: 2em;
    text-decoration: none;
    user-select: none;
    transition: color .3s ease-in-out;
}

.title-area > a:hover{
    color: #10b981;
    text-shadow: 0px 0px .2px #10b981;
}

.title-area > a::after{/*文章标题下划线动效*/
    content: "";
    position: absolute;
    width: 0%;
    height: 3px;
    top: 35px;
    left: 50%;
    right: 50%;
    color: #10b981;
    background-color: black;
    transition: all .3s ease-in-out;
}

.title-area > a:hover::after{
    left: 0%;
    right: 0%;
    width: 100%;
}

/*文章简介信息演示设置*/
.info-area{
    display: inline;
    /* min-width: 365px; */
    margin: 10px auto 10px auto;
}

.info-area > i{ /*简介内容*/
    color: rgb(130, 130, 130);
    font-style: normal;
}

.info-area > .divider{ /*简介分割线*/
    margin: 0 8px 0 8px;
    color: black;
    user-select: none;
}

/*简介信息前置图标*/
.info-area > .date::before{
    font-family: 'FontAwesome';
    content: "\f133";
    margin-right: 5px;
}

.info-area > .tag::before{
    font-family: 'FontAwesome';
    content: "\f114";
    margin-right: 5px;
}

.info-area > .word-count::before{
    font-family: 'FontAwesome';
    content: "\f1c2";
    margin-right: 5px;
}

.info-area > .read-time::before{
    font-family: 'FontAwesome';
    content: "\f017";
    margin-right: 5px;
}

.content-area{
    max-width: 500px;
    text-indent: 2em;
    text-align: left;
    margin: 10px auto 0 auto;
}

/*针对移动窄屏设备的响应式设计*/
@media screen and (max-width: 750px) {
    .avatar-area{
        display: none;
    }
}
</style>