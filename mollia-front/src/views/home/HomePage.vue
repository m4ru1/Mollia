<template>
    <div class="home-page">
        <div class="left-area">
            <ArticleList class="article-list"></ArticleList>
            <PaginationBoard
                :pager-count="7"
                :per-page="10"
                :total-number="1000">
            </PaginationBoard>
        </div>
        <div class="right-area">
            <TagBoard class="tag-board"></TagBoard>
            <ProfileBoard class="profile-board"></ProfileBoard>
        </div>
    </div>
</template>
    
<script setup>
import ProfileBoard from '../../components/profile/ProfileBoard.vue';
import ArticleList from '../../components/article/ArticleList.vue';
import TagBoard from '../../components/tag/TagBoard.vue';
import PaginationBoard from '../../components/pagination/PaginationBoard.vue';
</script>

<style>

.home-page{
    display: flex;
    justify-content: center;
    & > * {
        display: flex;
        flex-direction: column;
        margin: 30px;
    }
    .left-area{
        flex-grow: 0;
        flex-basis: 45vw;
        flex-shrink: 0;
    }
    .right-area{
        flex-grow: 0;
        flex-basis: 17vw;
        flex-shrink: 0;
        transition: all .5s ease-in-out;
        .tag-board {
            margin: 0 auto 0 auto;
            max-height: 350px;
        }
        .profile-board {
            margin: 30px 0 0 0;
            position: sticky;
            top: 65px;
        }
    }
}

/*针对移动窄屏设备的响应式设计*/
@media screen and (max-width: 800px){
    .home-page{
        flex-direction: column;
        & > *{
            flex-grow: 1;
            max-width: 100%;
            margin-left: 10px;
            margin-right: 10px;
        }
    }
    .tag-board{
        margin-right: 50px;
    }
}

@media screen and (max-width: 1300px) and (min-width: 800px){
    .home-page{
        flex-direction: column;
        align-items: center;
        & > *{
            flex-grow: 1;
            max-width: 80vw;
            margin-left: 10px;
            margin-right: 10px;
        }
        .right-area{
            width: 100%;
            flex-direction: row;
            .tag-board{
                width: 50%;
                margin: 0 10px 0 0;
            }
            .profile-board{
                width: 50%;
                margin: 0 0 0 10px;
            }
        }
    }
}

@media screen and (max-width: 1500px) and (min-width: 1300px){
    .home-page .right-area{
        flex-grow: 1;
        max-width: 25vw;
        *{
            flex-grow: 0;
        }
        .profile-board{
            flex-grow: 0;
        }
    }
    .home-page > .left-area{
        flex-grow: 1;
        max-width: 50vw;
        margin-left: 10px;
        margin-right: 10px;
    }
}
</style>