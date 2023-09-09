<template>
<div class="profile-option">
    <PageContent class="page-content">
        <template #page-title>
            <h2>网站信息</h2>
        </template>
        <template #page-view>
            <div class="switch-bar">
                <el-menu mode="horizontal" :ellipsis="false" @select="switchHandle">
                    <el-menu-item index="1">基本设置</el-menu-item>
                    <el-menu-item index="2">文章设置</el-menu-item>
                    <el-menu-item index="3">个人设置</el-menu-item>
                </el-menu>
            </div>
            <el-form v-if="currFormIndex === '1'" label-width="100px" class="option-form">
                <el-form-item label="网站Logo">
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
                <el-form-item label="网站域名">
                    <el-input v-model="domainName" style="width: 200px;"></el-input>
                </el-form-item>
                <el-form-item label="网站名称">
                    <el-input v-model="websiteName" style="width: 200px;"></el-input>
                </el-form-item>
                <el-form-item label="网站备案号">
                    <el-input v-model="websiteRecord" style="width: 200px;"></el-input>
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" @click="saveHandle('1')">保存</el-button>
                    <el-button type="danger" @click="initHandle">重置</el-button>
                </el-form-item>
            </el-form>
            <el-form v-else-if="currFormIndex === '2'" label-width="110px" class="option-form">
                <el-form-item label="文章默认封面图">
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
                    <el-button type="primary" @click="saveHandle('2')">保存</el-button>
                    <el-button type="danger" @click="initHandle">重置</el-button>
                </el-form-item>
            </el-form>
            <el-form v-else label-width="100px" class="option-form">
                <el-form-item label="博主名称">
                    <el-input v-model="name" style="width: 200px;"></el-input>
                </el-form-item>
                <el-form-item label="博主头像">
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
                <el-form-item label="博主座右铭">
                    <el-input v-model="motto" style="width: 200px;"></el-input>
                </el-form-item>
                <el-form-item label="Github主页">
                    <el-input v-model="githubUrl" style="width: 200px;"></el-input>
                </el-form-item>
                <el-form-item label="Twitter主页">
                    <el-input v-model="twitterUrl" style="width: 200px;"></el-input>
                </el-form-item>
                <el-form-item label="Facebook主页">
                    <el-input v-model="facebookUrl" style="width: 200px;"></el-input>
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" @click="saveHandle('2')">保存</el-button>
                    <el-button type="danger" @click="initHandle">重置</el-button>
                </el-form-item>
            </el-form>
        </template>
    </PageContent>
</div>
</template>

<script setup>
import { ref } from 'vue';
import { Plus } from '@element-plus/icons-vue'
import PageContent from '../../components/layout/PageContent.vue';

/****************************** 
        网站信息设置
*******************************/

// 基本设置
let logoUrl = ref('');      // 网站Logo
let domainName = ref('');   // 网站域名
let websiteName = ref('');  // 网站名称
let websiteRecord = ref('');// 网站备案号

// 个人设置
let avatarUrl = ref('');    // 博主头像
let name = ref('');         // 博主名称
let motto = ref('');        // 博主座右铭
let githubUrl = ref('');    // github主页
let twitterUrl = ref('');   // twitter主页
let facebookUrl = ref('');  // facebook主页

// 文章设置
let articleDefaultAvatarUrl = ref('');  // 文章默认封面图

/****************************** 
        表单逻辑
*******************************/

let currFormIndex = ref('1');
function switchHandle(index){ // 表单切换逻辑
    currFormIndex.value = index;
}

function initHandle(){  // 表单恢复默认值（即重新拉取并填充当前值）

}

function saveHandle(index){  // 保存配置

}
</script>

<style scoped>
.profile-option{
    display: flex;
    background-color: var(--theme-pagecontent-bg-color);
    user-select: none;
    .switch-bar{
        padding: 15px;
    }
    .option-form{
        padding: 15px 15px 15px 20px;
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
</style>