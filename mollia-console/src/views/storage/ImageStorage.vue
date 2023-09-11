<template>
    <div class="image-option">
        <PageContent class="page-content">
            <template #page-title>
                <h2> 图床管理 </h2>
            </template>
            <template #page-view>
                <div class="search-bar-box">
                    <div class="search-bar">
                        <span> 图片名: </span>
                        <el-input 
                            v-model="tempSearchPhotoName"
                            style="width: 200px; margin-right: 20px;"
                            placeholder="请输入文章标题" clearable>
                        </el-input>
                        <el-button type="primary" @click="searchHandle">
                            检索
                        </el-button>
                        <el-button type="success" @click="openUploadBoard">
                            上传
                        </el-button>
                    </div>
                </div>
                <div class="photo-table">
                    <el-card 
                        class="photo-card"
                        v-for="(item, index) in activePhotoData"
                        :body-style="{ padding: '0px' }"
                        :key="index"
                        @click="openOptBoard(item.name, item.url)">
                        <div class="img-view">
                            <img :src="item.url">
                        </div>
                        <div class="img-name">
                            <span> {{ item.name }} </span>
                        </div>
                    </el-card>
                </div>
            </template>
        </PageContent>
        <div :class="{'upload-board': true, 'upload-board-active': isUploadBoardVisable}">
            <el-form class="upload-form" label-width="120px">
                <h2>图像上传</h2>
                <el-form-item label="图像文件名" >
                    <el-input placeholder="请输入文件图像名" style="width: 200px"></el-input>
                </el-form-item>
                <el-form-item label="图像上传区">
                    <el-upload
                        class="avatar-uploader"
                        action="javascript(0);"
                        :show-file-list="false"
                        :on-success="handleAvatarSuccess"
                        :before-upload="beforeAvatarUpload">
                        <img v-if="imgUrl" :src="imgUrl" class="avatar"/>
                        <el-icon v-else class="avatar-uploader-icon"><Plus /></el-icon>
                    </el-upload>
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" @click="uploadHandle">
                        上传
                    </el-button>
                    <el-button type="danger" @click="closeUploadBoard">
                        取消
                    </el-button>
                </el-form-item>
            </el-form>
        </div>
        <div :class="{'opt-board': true, 'opt-board-active': isOptBoardVisable}">
            <img :src="currViewPhoto.url" class="big-img" @click="optBoardKeyHandle">
            <div class="opt-bar">
                <ul>
                    <li @click="optZoomoutHandle"><el-icon color="silver" size="25"><ZoomOut /></el-icon></li>
                    <li @click="optZoominHandle"><el-icon color="silver" size="25"><ZoomIn /></el-icon></li>
                    <li @click="optDownloadHandle"><el-icon color="silver" size="25"><Download /></el-icon></li>
                    <li @click="optDeleteHandle"><el-icon color="silver" size="25"><Delete /></el-icon></li>
                </ul>
            </div>
        </div>
    </div>
</template>

<script setup>
import PageContent from '../../components/layout/PageContent.vue';
import { Plus, ZoomOut, ZoomIn, Delete, Download,  } from '@element-plus/icons-vue';
import { ref, computed } from 'vue';

// 图片搜索逻辑
let photoData = ref([]);
let tempSearchPhotoName = ref('');
let searchPhotoName = ref('');
let activePhotoData = computed(() => {
    return photoData.value.filter((item) => item.name.includes(searchPhotoName.value));
});

function searchHandle(){
    searchPhotoName.value = tempSearchPhotoName.value;
}

// 上传逻辑
let isUploadBoardVisable = ref(false);  // 上传面板是否打开
let uploadImageName = ref('');            // 上传文件名

function openUploadBoard(){ // 打开上传面板
    isUploadBoardVisable.value = true;
}

function closeUploadBoard(){ // 关闭上传面板
    isUploadBoardVisable.value = false;
}

function uploadHandle(){ // 文件上传逻辑

}

// 图片操作逻辑
let isOptBoardVisable = ref(false);         // 图片操作面板是否打开
let currViewPhoto = ref({name: '', url: ''});// 当前展示图像数据

function openOptBoard(name, url){ // 打开图片操作面板
    isOptBoardVisable.value = true;
    currViewPhoto.value.name = name;
    currViewPhoto.value.url = url;
    console.log(currViewPhoto.value);
}

function closeOptBoard(){         // 关闭图片操作面板

}

function optBoardKeyHandle(event){     // 图片面板事件监听
    isOptBoardVisable.value = false;
}

function optZoomoutHandle(){ // 图片缩小

}

function optZoominHandle(){ // 图片放大

}

function optDownloadHandle(){ // 图片下载

}

function optDeleteHandle(){ // 图片删除

}

// 测试数据
for(let i = 0; i < 10; i++){
    photoData.value.push({
        name: 'girl1111.jpg',
        url: 'https://img1.imgtp.com/2023/09/10/r4LIDUqB.png'
    });
}

</script>

<style scoped>
.image-option{
    display: flex;
    overflow: hidden;
    position: relative;
    background-color: var(--theme-pagecontent-bg-color);
    .search-bar-box{
        padding: 15px;
        user-select: none;
        .search-bar{
            padding: 10px 0px 20px 10px;
            border-bottom: 1px solid #ccc;
            & > span{
                margin: auto 0px;
            }
        }
    }
    .photo-table{
        display: flex;
        justify-content: center;
        flex-wrap: wrap;
        padding: 10px;
        transition: all .2s ease-out;
        .photo-card{
            margin: 5px;
            transition: all .1s ease-out;
            box-shadow: 0px 0px 10px #ccc;
            overflow: hidden;
            user-select: none;
            cursor: pointer;
            .img-view{
                text-align: center;
                padding: 5px;
                transition: all .1s ease-out;
                & > img{
                    width: auto;
                    height: 200px;
                    transition: all .1s ease-out;
                }
            }
            .img-name{
                max-width: 200px;
                font-size: 16px;
                padding-bottom: 10px;
                overflow: hidden;
                text-align: center;
            }
        }
    }
    .upload-board{
        display: none;
        position: absolute;
        width: 100%;
        height: 100%;
        text-align: center;
        background-color: rgba(0, 0, 0, .3);
        user-select: none;
        .upload-form{
            margin: 50px auto;
            width: 500px;
            padding: 20px;
            border-radius: 20px;
            background-color: white;
            .avatar-uploader{
                width: 178px;
                height: 178px;
                display: block;
                border: 1px dashed black;
                border-radius: 6px;
                cursor: pointer;
                transition: all .3s ease-out;
                overflow: hidden;
                position: relative;
                &:hover{
                    border-color: var(--theme-primary-color);
                }
                .avatar-uploader-icon{
                    font-size: 28px;
                    color: #8c939d;
                    width: 178px;
                    height: 178px;
                    text-align: center;
                }
            }
        }
    }
    .upload-board-active{
        display: block;
    }
    .opt-board{
        width: 100%;
        height: 100%;
        display: none;
        position: absolute;
        text-align: center;
        user-select: none;
        -webkit-user-drag: none;
        background-color: rgba(0, 0, 0, .8); /*TODO: 覆盖全屏考虑全局蒙版，用pinia变量控制开关 */
        .big-img{
            margin-top: 50px;
            padding: 20px;
            max-width: 30vw;
            height: auto;
            max-height: 80vh;
            cursor: pointer;
        }
        .opt-bar{
            display: flex;
            justify-content: center;
            background-color: transparent;
            & > ul{
                display: flex;
                flex-direction: row;
                list-style: none;
                & > li{
                    margin: 0px 5px;
                    & > i > svg:hover{
                        cursor: pointer;
                        fill: red;
                    }
                }
            }
        }
    }
    .opt-board-active{
        display: block;
    }
}
</style>