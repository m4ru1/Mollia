<template>
    <div class="article-page">
        <div class="left-area">
            <ArticleTimeline></ArticleTimeline>
        </div>
        <div class="right-area">
            <TagBoard></TagBoard>
            <ProfileBoard></ProfileBoard>
        </div>
    </div>
</template>
    
<script setup>
import ArticleTimeline from '../../components/article/ArticleTimeline.vue';
import TagBoard from '../../components/tag/TagBoard.vue';
import ProfileBoard from '../../components/profile/ProfileBoard.vue';

</script>

<style scoped>
.article-page{
    display: flex;
    justify-content: center;
    & > .left-area, & > .right-area {
        display: flex;
        flex-direction: column;
        margin: 30px;
    }
    .left-area{
        display: flex;
        flex-grow: 0;
        flex-basis: 35vw;
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
    .article-page{
        flex-direction: column;
        & > .left-area, & > .right-area {
            flex-grow: 1;
            max-width: 100%;
            margin-left: 15px;
            margin-right: 15px;
        }
    }
    .tag-board{
        margin-right: 50px;
    }
}

@media screen and (max-width: 1300px) and (min-width: 800px){
    .article-page{
        flex-direction: column;
        align-items: center;
        & > .left-area, & > .right-area {
            flex-grow: 1;
            max-width: 80vw;
            margin-left: 10px;
            margin-right: 10px;
        }
        .left-area{
            width: 100%;
            flex-grow: 1;
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
</style>