<template>
    <div class="tag-board">
        <div class="tag-header">
            <p>
                文章分类
            </p>
        </div>
        <div class="tag-content">
            <a class="tag-item" v-for="item in testData" :key="item.id" :href="`https:localhost`">
                <i style="font-family:FontAwesome;margin-right: 8px;" class="fa-thin fa-tag"></i>
                <span>{{ item.name }}</span>
                <span>（{{ item.articleCount }}）</span>
            </a>
        </div>
    </div>
</template>

<script setup>

const testData = [
    {name:'C', id:751, articleCount: 15},
    {name:'C++', id:752, articleCount: 16},
    {name:'Python', id:753, articleCount: 17},
    {name:'Java', id:754, articleCount: 18},
    {name:'Golang', id:755, articleCount: 19},
    {name:'Rust', id:756, articleCount: 12},
    {name:'JavaScript', id:757, articleCount: 13},
    {name:'Lua', id:758, articleCount: 14},
    {name:'Mojo', id:759, articleCount: 17},
    {name:'Mojo1', id:760, articleCount: 17},
    {name:'Mojo2', id:761, articleCount: 17},
    {name:'Mojo3', id:762, articleCount: 17},
]
</script>

<style>
.tag-board{
    display: flex;
    flex-direction: column;
    flex-shrink: 0;
    width: 100%;
    position: relative;
    background-color: #111827;
    opacity: .7;
    border-radius: 20px;
    color: white;
    overflow: hidden;
}

.tag-header{
    flex-grow: 1;
    text-align: center;
    padding: 32px 24px 24px 24px;
    margin-bottom: 10px;
    border-radius: 20px;
    box-shadow: 0 0 1.1px white;
    font-size: 24px;
    font-weight: 700;
}

.tag-header > p {
    margin: 0;
}

.tag-content{
    display: flex;
    flex-direction: column;
    overflow: auto;
    padding-bottom: 10px;
}

/*滚动条美化，非标准语法，Firefox PC/Mobile以及Safari均不支持*/
.tag-content::-webkit-scrollbar{ /*滚动条宽度*/
    width: 2px;
    height: 1px;
}

.tag-content::-webkit-scrollbar-track{ /*滚动条轨道*/
    background-color: transparent;
    border-radius: 2em;
}

.tag-content::-webkit-scrollbar-thumb{ /*滚动条滑块*/
    background-color: rgb(255, 255, 255, .7);
    border-radius: 2em;
}

.tag-content > .tag-item{
    margin: 10px auto 0 20px;
    font-size: 17px;
    color: white;
    text-decoration: none;
    cursor: pointer;
    transition: color .1s ease;
}

.tag-content > .tag-item:hover{
    color: #10b981;
}
</style>