<template>
    <div class="article-category">
        <PageContent class="page-content">
            <template #page-title>
                <h2>文章分类</h2>                                   
            </template>
            <template #page-view>
                <div class="search-bar-border">
                    <div class="search-bar">
                        <span class="search-bar-input-tag">
                            类别名
                        </span>
                        <el-input class="search-bar-input"
                            placeholder="请输入类别名"
                            v-model="searchNameTemp">
                        </el-input>
                        <el-button class="search-bar-btn"
                            type="info"
                            @click="searchHandle">
                            搜索
                        </el-button>
                        <el-button class="open-create-btn"
                            type="primary"
                            @click="openCreateHandle">
                            创建类别
                        </el-button>
                    </div>
                </div>
                <div class="category-table">
                    <el-table v-if="selectedCategoryData.length" 
                        :data="selectedCategoryData" border style="width: 100;">
                        <el-table-column prop="name" label="类别名"></el-table-column>
                        <el-table-column prop="count" label="文章数"></el-table-column>
                        <el-table-column prop="createDate" label="创建日期"></el-table-column>
                        <el-table-column prop="updateDate" label="更新日期"></el-table-column>
                        <el-table-column label="操作">
                            <template v-slot="{ row }">
                                <el-button @click="openEditHandle(row)" type="primary">编辑</el-button>
                                <el-button @click="deleteHandle(row)" type="danger">删除</el-button>
                            </template>
                        </el-table-column>
                    </el-table>
                    <el-empty v-else description="这里空空如也"></el-empty>
                </div>
                <div class="pagination-box" v-if="selectedCategoryData.length">
                    <el-pagination background layout="prev, pager, next" :total="selectedCategoryData.length"></el-pagination>
                </div>
            </template>
        </PageContent>
        <div :class="{'create-board': true, 'create-board-active': isCreatingTag}">
            <el-form label-width="100px" class="create-form">
                <h2 class="create-form-header">创建类别</h2>
                <el-form-item label="类别名称">
                    <el-input v-model="tagName" placeholder="请输入类别名称" style="width: 200px;"></el-input>
                </el-form-item>
                <el-form-item label="类别缩略图">
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
                    <div class="op-btn-board">
                        <el-button @click="createHandle" type="primary">创建</el-button>
                        <el-button @click="cancelCreateHandle">取消</el-button>
                    </div>
                </el-form-item>
            </el-form>
        </div>
    </div>
</template>

<script setup>
import PageContent from '../../components/layout/PageContent.vue';
import { ref, computed } from 'vue';
import { Plus } from '@element-plus/icons-vue';

// 数据存储
let categoryData = ref([]);
let searchNameTemp = ref('');
let searchName = ref('');
let selectedCategoryData = computed(()=>{
    return categoryData.value.filter((item) => item.name.includes(searchName.value));
});

// 搜索逻辑
function searchHandle(){
    searchName.value = searchNameTemp.value;
}

// 创建逻辑
let isCreatingTag = ref(false);
let tagName = ref('');

function openCreateHandle(){    // 打开创建面板
    isCreatingTag.value = true;
}

function createHandle(){    // 类别创建逻辑
    
}

function cancelCreateHandle(){  // 关闭创建面板
    isCreatingTag.value = false;
}

/***
    表格操作逻辑 
***/

// 编辑逻辑
function openEditHandle(){  // 打开编辑面板

}

function editHandle(){  // 类别编辑逻辑

}

// 删除逻辑
function deleteHandle(){  // 删除逻辑

}

// 测试数据
for (let i = 0; i < 100; i++){
    categoryData.value.push({
        name: `类别${i+1}`,
        count: `${i*2}`,
        createDate: `2023-9-${i}`,
        updateDate: `2023-9-${i}`
    });
}

</script>

<style scoped>
.article-category{
    display: flex;
    overflow: hidden;
    position: relative;
    background-color: var(--theme-pagecontent-bg-color);
    .search-bar-border{
        padding: 15px;
        .search-bar{
            display: flex;
            padding: 15px 15px 15px 0px;
            box-sizing: border-box;
            border-bottom: 1px solid #ccc;
            user-select: none;
            .search-bar-input-tag{
                margin: auto 10px auto 0px;
            }
            .search-bar-input{
                width: 200px;
                margin-right: 15px;
            }
            .open-create-btn{
                margin-left: auto;
            }
        }
    }
    .category-table{
        padding: 0px 15px 30px 15px;
        max-width: 90vw;
        min-width: 1000px;
    }
    .pagination-box{
        display: flex;
        justify-content: end;
        padding: 5px 30px 15px 15px;
    }
    .create-board{
        display: none;
        position: absolute;
        width: 100%;
        height: 100%;
        background-color: rgba(60, 60, 60, .7);
        user-select: none;
        .create-form{
            margin: 5vh auto;
            padding: 100px 40px;
            background-color: white;
            width: 600px;
            border-radius: 20px;
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
            .create-form-header{
                padding-bottom: 30px;
            }
        }
        z-index: 20;
    }
    .create-board-active{
        display: block;
    }
}
</style>