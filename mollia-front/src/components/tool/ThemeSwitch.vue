<template>
<div class="theme-switch">
    <div class="switch-bar">
        <input id="switch" type="checkbox">
        <label for="switch"></label>
    </div>
</div>
</template>

<script setup>

</script>

<style>

.theme-switch{
    display: flex;
}

.switch-bar{
    width: 50px;
    height: 20px;
    margin: auto 0 auto 0;
}

.switch-bar > input{
    display: none;
}

.switch-bar > label{
    display: flex;
    width: 100%;
    height: 20px;
    background-color: #222;
    position: relative;
    cursor: pointer;
    user-select: none;
    border-radius: 90px;
    box-shadow: 0 0 2px #ccc;
}

.switch-bar > label::before{
    content: "";
    width: 100%;
    height: 100%;
    border-radius: inherit;
    opacity: 1;
    transition: all .4s ease;
    background: linear-gradient(330deg, #ccc 30%, #10b981 100%);
}

.switch-bar > label::after{ /*左右滑动的小钮扭*/
    content: "";
    position: absolute;
    left: 0px;
    height: 20px;
    width: 20px;
    box-shadow: 0 0 1px #222;
    border-radius: 50%;
    /*日间模式图标*/
    background-color: #ccc;
    background-image: url("../../assets/light-theme.svg");
    background-size: 14px;
    background-repeat: no-repeat;
    background-position: 3px 48%;
    transition: all .4s ease;
}

.switch-bar > input:checked ~ label::before{ /*暗夜模式：滑轨*/
    background-color: #222;
    opacity: 0;
}

.switch-bar > input:checked ~ label::after{ /*暗夜模式：小扭扭*/
    left: 30px;
    background-color: #444;
    /*暗夜模式图标*/
    background-image: url("../../assets/dark-theme.svg");
    background-size: 14px;
    background-repeat: no-repeat;
    background-position: 3px 48%;
}
</style>