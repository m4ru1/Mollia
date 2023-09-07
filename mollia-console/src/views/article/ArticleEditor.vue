<template>
    <div class="article-editor">
        <PageContent class="page-content">
            <template #page-view>
                <div class="editor-header">
                    <el-input class="editor-title-input" v-model="articleTitle" placeholder="请输入文章标题"></el-input>
                    <el-button type="primary" class="editor-btn" @click="postHandle">
                        发布文章
                    </el-button>
                    <el-button type="success" class="editor-btn" @click="draftHandle">
                        保存草稿
                    </el-button>
                </div>
                <div id="editor"></div>
            </template>
        </PageContent>
        <div :class="{'post-board':true, 'post-board-active': isPostBoardActive}">
            <el-form class="post-form" label-width="100px">
                <h2 class="post-form-title">文章发布</h2>
                <el-form-item label="文章标签">
                    <el-tag v-for="(item, index) in articleTags"
                        :key="index"
                        class="articleTag"
                        closable
                        @close="tagCloseHandle">
                        {{ item.name }}
                    </el-tag>
                    <el-autocomplete v-if="isTagEditing" 
                        placeholder="请输入标签名称"
                        v-model="currEditTag"
                        :fetch-suggestions="tagAutoFetch"
                        @keydown="tagAddHandle"
                        @blur="tagLoseFocus"
                        autofocus> 

                    </el-autocomplete>
                    <el-button v-else :icon="Plus" @click="tagEditHandle"></el-button>
                </el-form-item>
                <el-form-item label="文章分类">
                    <el-select class="input" placeholder="请选择文章类别" v-model="articleCategory">
                        <el-option 
                            v-for="(item, index) in articleCategoryChoices"
                            :key="index"
                            :label="item.label"
                            :value="item.value"
                        />
                    </el-select>
                </el-form-item>
                <el-form-item label="文章缩略图">
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
                <el-form-item label="置顶">
                    <el-radio-group v-model="isArticleTop" class="input">
                        <el-radio :label="true">打开</el-radio>
                        <el-radio :label="false">关闭</el-radio>
                    </el-radio-group>
                </el-form-item>
                <el-form-item label="发布形式">
                    <el-radio-group v-model="isArticlePublic" class="input">
                        <el-radio :label="true">公开</el-radio>
                        <el-radio :label="false">私密</el-radio>
                    </el-radio-group>
                </el-form-item>
                <el-form-item label="发布日期">
                    <el-date-picker 
                        v-model="articlePublishDate"
                    />
                </el-form-item>
                <el-form-item>
                    <div class="op-btn-board">
                        <el-button type="primary">发布</el-button>
                        <el-button @click="cancelHandle">取消</el-button>
                    </div>
                </el-form-item>
            </el-form>
        </div>        
    </div>
</template>

<script setup>
import { Plus } from '@element-plus/icons-vue';
import { onMounted, computed, ref, shallowRef, onUnmounted } from 'vue';
import PageContent from '../../components/layout/PageContent.vue';
import Cherry from 'cherry-markdown';
import 'cherry-markdown/dist/cherry-markdown.css';

// 创建编辑器实例
const config = {
    id: 'editor',
    value: '',
    callback: {
        afterchange: function(md, html){
            console.log('change');
        }
    }
};

let cherry;

onMounted(() => {
    cherry = new Cherry(config);
});

// 文章数据存储
let articleTitle = ref(''); // 文章标题
let articleData = ()=>{     // 文章内容
    return cherry.getMarkdown();
};
let articleTags = ref([]);  // 文章标签
let articleTagChoices = ref([]);  // 文章现有标签
let currEditTag = ref('');  // 当前编辑标签
let isTagEditing = ref(false); // 当前是否正在编辑标签
let articleCategoryChoices = ref([]); // 文章现有类别
let articleCategory = ref('');  // 文章类别
let isArticleTop = ref(false);  // 文章是否置顶
let articlePublishDate = ref('');  // 文章发布日期
let isArticlePublic = ref(true);  // 是否公开

let isPostBoardActive = ref(false);

// 文章发布逻辑
function postHandle(){ // 发布逻辑
    isPostBoardActive.value = true;
}

function draftHandle(){ // 草稿逻辑

}

function cancelHandle(){ // PostBoard取消逻辑
    isPostBoardActive.value = false;
}

// 表单操作逻辑
function tagCloseHandle(args){  // 删除Tag
    console.log(articleData());
}

function tagEditHandle(){  // 编辑Tag
    isTagEditing.value = true;
}

function tagLoseFocus(){  // 标签编辑框失去焦点
    currEditTag.value = '';
    isTagEditing.value = false;
}

function tagAddHandle(event){ // 添加Tag
    if(event.key === 'Enter'){

    }else if(event.key === 'Escape'){
        tagLoseFocus();
    }
}

function tagAutoFetch(queryString){  // 自动搜索匹配现存Tag

}

function beforeAvatarUpload(){ // 上传文章封面图准备逻辑

}

function handleAvatarSuccess(){ // 上传文章封面图成功逻辑

}


</script>

<style scoped>
.article-editor{
    display: flex;
    overflow: hidden;
    position: relative;
    background-color: var(--theme-pagecontent-bg-color);
    .page-content{
        display: flex;
        flex-direction: column;
        border-radius: 20px;
        .editor-header{
            display: flex;
            padding:20px 15px 15px 15px;
            .editor-title-input{
                margin-right: 100px;
            }
        }
        #editor{
            max-width: 2600px;
            max-height: 70vh;
            padding: 5px 15px 10px 15px;
        }
    }
    .post-board{
        display: none;
        position: absolute;
        height: 100%;
        width: 100%;
        background-color: rgb(60, 60, 60, .7);
        user-select: none;
        .post-form{
            margin: 5vh auto;
            background-color: white;
            padding: 100px 40px;
            width: 600px;
            border-radius: 20px;
            .post-form-title{
                padding-bottom: 30px;
            }
            .form-item{
                .input{
                    display: inline;
                    width: 300px;
                }
            }
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
            .op-btn{
                display: flex;
                justify-content: end;
                align-items: end;
            }
        }
        z-index: 20;
    }
    .post-board-active{
        display: block;
    }
}
</style>