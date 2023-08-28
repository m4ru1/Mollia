<template>
    <div class="profile-board">
        <img :src="avatorURL">
        <div class="name-bar">
            {{ name }}
        </div>
        <div class="motto-bar" v-html="motto">
        </div>
        <div class="audit-bar">
            <a class="audit-item article-count">{{ articleCount }}</a>
            <a class="audit-item tag-count">{{ tagCount }}</a>
            <a class="audit-item visitor-count">{{ visitorCount }}</a>
        </div>
        <div class="link-bar">
            <a class="link-item" :href="githubLink">
                <i class="fa-brands fa-github"></i>
            </a>
            <a class="link-item" :href="twitterLink">
                <i class="fa-brands fa-twitter"></i>
            </a>
            <a class="link-item" :href="facebookLink">
                <i class="fa-brands fa-facebook"></i>
            </a>
        </div>
    </div>
</template>

<script setup>

// profile
let avatorURL = "https://img1.imgtp.com/2023/08/14/6jg4sgNZ.png";
let name =  "Mollia";
let motto = " 三年羁旅长堪醉，两盏粗酒饮狂风。";
// links
let githubLink = "https://github.com/m4ru1";
let twitterLink = "https://mollia.space";
let facebookLink = "https://mollia.space";
// audit
let articleCount = 12;
let tagCount = 12;
let visitorCount = 256;
</script>

<style>
.profile-board{
    display: flex;
    flex-direction: column;
    justify-content: center;
    background-color: #111827;
    opacity: .7;
    border-radius: 20px;
    overflow: hidden;
    color: white;
    padding: 0 30px 0 30px;
    user-select: none;
}

.profile-board > img{
    border-radius: 90px;
    width: 126px;
    height: 126px;
    margin: 25px auto 10px auto;
    user-select: none;
    -webkit-user-drag: none;
}

.profile-board > .name-bar{
    display: flex;
    font-size: 24px;
    justify-content: center;
}

.profile-board > .motto-bar{
    display: flex;
    justify-content: center;
}

.profile-board > .audit-bar{
    display: flex;
    justify-content: center;
    padding-left: 10%;
    padding-right: 10%;
}

.audit-bar > .audit-item{
    flex-grow: 1;
    position: relative;
    display: block;
    text-align: center;
    margin: 15px 0px 30px 0px;
}

.audit-bar > .audit-item::after{
    position: absolute;
    width: 100%;
    top: 100%;
    left: 0%;
    right: 0%;
}

.audit-bar > .article-count::after{
    content: "Articles";
}

.audit-bar > .tag-count::after{
    content: "Tags";
}

.audit-bar > .visitor-count::after{
    content: "Visitors";
}

.profile-board > .link-bar{
    display: flex;
    justify-content: center;
    font-family: 'FontAwesome';
    margin: 10px auto 25px auto;
}

.profile-board > .link-bar > .link-item{
    font-style: normal;
    font-size: 20px;
    color: white;
    margin: 0 10px 0 10px;
    text-decoration: none;
    & i{
        font-style: normal;
    }
}
</style>