<template>
    <div class="nav-bar">
        <strong id="logo-text">  Mollia  </strong>
        <input class="menu-btn" id="menu-btn" type="checkbox">
        <label for="menu-btn" class="menu-icon">
            <span class="nav-icon"></span>
        </label>
        <ThemeSwitch id="theme-switch"></ThemeSwitch>
        <SearchBar id="search-bar"></SearchBar>
        <div class="menu">
            <router-link class="menu-item nav-link" to="/">主页</router-link>
            <router-link class="menu-item nav-link" to="/article">文章</router-link>
            <router-link class="menu-item nav-link" to="/tag">专栏</router-link>
            <router-link class="menu-item nav-link" to="/about">关于</router-link>
        </div>
    </div>
</template>

<script setup>
import SearchBar from "../tool/SearchBar.vue"
import ThemeSwitch from "../tool/ThemeSwitch.vue";
// 搜索逻辑

// 搜索时主界面蒙版效果

// 黑夜模式逻辑

</script>

<style>

/* header */
.nav-bar{
    position: fixed;
    width: 100%;
    background-color: #1b1b1b;
    z-index: 3;
}

#logo-text{
    display: block;
    padding: 10px 20px;
    float: left;
    font-size: 30px;
    color: #10b981;
    text-shadow: 0 0 3px pink;
    user-select: none;
}

.menu-item{
    color: white;
    display: block;
    font-size: 16px;
    padding: 20px 20px;
    text-decoration: none;
    user-select: none;
}

.menu-item:hover{
    color: black;
    font-weight: 700;
    background-color: #10b981;
}

#search-bar{
    display: block;
    float: right;
    padding: 12px 20px 15px 0;
}

#theme-switch{
    display: block;
    float: right;
    padding: 18px 0px;
}

/* menu */
.menu{
    clear: both;
    max-height: 0;  /*缩起菜单*/
    transition: max-height ease-in-out .2s;
    overflow: hidden;
}

/* menu icon */
.nav-bar > .menu-icon{
    cursor: pointer;
    display: block;
    float: right;
    padding: 28px 20px;
    position: relative;
    user-select: none;
}

.nav-bar .menu-icon .nav-icon{
    background-color: #ddd;
    display: block;
    height: 2px;
    position: relative;
    transition: background-color .2s ease-out;
    width: 25px;
}

.nav-bar .menu-icon .nav-icon::before,
.nav-bar .menu-icon .nav-icon::after{
    background-color: #ddd;
    content: '';
    display: block;
    height: 100%;
    position: absolute;
    transition: all .2s ease-out;
    width: 100%;
}

.nav-bar .menu-icon .nav-icon::before{
    top: 6px;
}

.nav-bar .menu-icon .nav-icon::after{
    top: -6px;
}

/* menu-btn */

#menu-btn{
    display: none;
}

.menu-btn:checked ~ .menu{
    max-height: 400px;
}

.menu-btn:checked ~ .menu-icon .nav-icon{
    background-color: transparent;
}

.menu-btn:checked ~ .menu-icon .nav-icon::before{ /*菜单按钮选中后的旋转动效*/
    top: 0;/* 从原先的上下横线分隔，到现在的交汇并呈X状 */
    transform: rotate(-45deg);
}

.menu-btn:checked ~ .menu-icon .nav-icon::after{
    top: 0;
    transform: rotate(45deg);
}

/* 宽度大于1000px的设备，将菜单按钮隐藏，并将菜单选项至于界面左端*/
@media screen and (min-width: 1000px){
    .nav-bar #theme-switch{
        margin-right: 50px;
    }
    .nav-bar .menu-icon, .nav-bar .nav-icon{
        display: none;
    }
    .menu{
        position: absolute;
        max-height: 100px;
        height: 100px;
        width: 500px;
        top: 20px;
        left: 125px;
        z-index: 1000;
        overflow: visible;
        .menu-item{
            display: inline;
        }
    }
}

/*
.nav-link{
    font-size: 20px;
    padding: 1em;
    color: white;
    text-decoration: none;
    transition: all 0.2s;
    user-select: none;
    -webkit-user-drag: none;
}

.nav-link:hover{
    font-weight: 10em;
    background-color: rgba(10, 160, 100, 10);
    box-shadow: 0 0 0.5px #10b981;
    border-radius: 30px;
    color: #000;
}

.nav-link:active{
    font-weight: 10em;
    background-color: rgba(20, 120, 90, 20);
    box-shadow: 0 0 0.5px #10b981;
    border-radius: 30px;
    color: #111;
}

.nav-bar > .left-area{
    display: flex;
    flex-basis: 500px;
    min-width: 500px;
    max-width: 42vw;
    flex-grow: 1;
    justify-content: left;
}

.nav-bar > .right-area{
    display: flex;
    flex-basis: 420px;
    min-width: 420px;
    max-width: 26vw;
    flex-grow: 1;
    justify-content: right; 此区域元素靠右显示，即从flex容器主轴终止线开始排列
}

.nav-bar > .profile-area{
    display: flex;
    flex-basis: 300px;
    max-width: 300px;
    justify-content: right;
    margin: auto 0 auto 0;
    user-select: none;
}

.nav-bar > .profile-area > img{
    height: 45px;
    width: 45px;
    margin-right: 15px;
    border-radius: 90px;
    transition: all .4s ease;
    -webkit-user-drag: none;
}

.nav-bar > .profile-area > img:hover{
    box-shadow: 0 0 10px #10b981;
}

.mobile-nav-bar, .mobile-area{
    display: none;
} */
/**/
</style>