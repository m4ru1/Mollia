<template>
    <div class="home">
        <div class="first-wrap">
            <div class="avatar-board">
            <img class="avatar" src="https://img1.imgtp.com/2023/08/14/6jg4sgNZ.png">
            <div class="info">
                <h1>{{ name }}</h1>
                <p>{{ signature }}</p>
            </div>
        </div>
        <div class="statistic-board">
            <div class="statistic-tag">
                <span class="statistic-item"> 文章数 
                    <span class="statistic-item-data">{{ articleNubmer }}</span>
                </span>
            </div>
            <span class="statistic-decollator"> | </span>
            <div class="statistic-tag">
                <span class="statistic-item"> 分类数
                    <span class="statistic-item-data">{{ categoryNumber }}</span>
                </span>
            </div>
            <span class="statistic-decollator"> | </span>
            <div class="statistic-tag">
                <span class="statistic-item"> 访客数
                    <span class="statistic-item-data">{{ vistorNumber }}</span>
                </span>
            </div>
        </div>
        </div>
        <div class="second-wrap">
            <div class="log-board">
            
            </div>
            <div class="message-board">

            </div>
        </div>
    </div>
</template>

<script setup>
import { ref } from 'vue';

let name = ref('Mollia');
let signature = ref('三年羁旅长堪醉，两盏粗酒饮狂风。');
let articleNubmer = ref('12');
let categoryNumber = ref('12');
let vistorNumber = ref('33608');

</script>

<style scoped>
.home{
    background-color: #ccc;
    .first-wrap{
        display: flex;
        .avatar-board, .statistic-board{
            display: flex;
            justify-content: center;
            flex-direction: row;
            flex: 1;
            background-color: white;
            padding: 20px;
            margin: 20px;
            border-radius: 20px;
            user-select: none;
        }
        .avatar-board{
            .avatar{
                width: 200px;
                height: 200px;
                padding: 10px;
                border-radius: 1000px;
            }
            .info{
                margin-left: 50px;
                margin-top: auto;
                margin-bottom: auto;
                width: 300px;
            }
        }
        .statistic-board{
            align-items: center;
            .statistic-tag{
                text-align: center;
                flex: 1;
                margin-bottom: 30px;
                .statistic-item{
                    display: block;
                    width: 100px;
                    font-size: 20px;
                    margin: 0 auto;
                    position: relative;
                    user-select: none;
                    .statistic-item-data{
                        width: 100px;
                        position: absolute;
                        top: 150%;
                        left: 0%;
                        right: 0%;
                    }
                }
                .statistic-decollator{
                    display: block;
                    padding: 0 50px;
                }
            }
        }
    }
}
</style>