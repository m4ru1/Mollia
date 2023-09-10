<template>
    <div class="theme-option">
        <PageContent class="page-content">
            <template #page-title>
                <h2> 主题设置 </h2>
            </template>
            <template #page-view>
                <el-form class="theme-form" label-width="110px">
                    <el-form-item label="白天模式背景图">
                        <el-upload
                            class="avatar-uploader"
                            action="javascript(0);"
                            :show-file-list="false"
                            :on-success="handleAvatarSuccess"
                            :before-upload="beforeAvatarUpload">
                            <img v-if="dayBgUrl" :src="imgUrl" class="avatar"/>
                            <el-icon v-else class="avatar-uploader-icon"><Plus /></el-icon>
                        </el-upload>
                    </el-form-item>
                    <el-form-item label="黑夜模式背景图">
                        <el-upload
                            class="avatar-uploader"
                            action="javascript(0);"
                            :show-file-list="false"
                            :on-success="handleAvatarSuccess"
                            :before-upload="beforeAvatarUpload">
                            <img v-if="nightBgUrl" :src="imgUrl" class="avatar"/>
                            <el-icon v-else class="avatar-uploader-icon"><Plus /></el-icon>
                        </el-upload>
                    </el-form-item>
                    <el-form-item>
                        <el-button type="primary" @click="saveHandle">保存</el-button>
                        <el-button type="danger" @click="initHandle">重置</el-button>
                    </el-form-item>
                </el-form>
            </template>
        </PageContent>
    </div>
</template>

<script setup>
import PageContent from '../../components/layout/PageContent.vue';
import { Plus } from '@element-plus/icons-vue'
import { ref } from 'vue';

let dayBgUrl = ref('');
let nightBgUrl = ref('');

</script>

<style scoped>
.theme-option{
    display: flex;
    overflow: hidden;
    position: relative;
    background-color: var(--theme-pagecontent-bg-color);
    .theme-form{
        padding: 35px 15px 15px 20px;
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