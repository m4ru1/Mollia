<template>
<div class="page-footer">
    <div class="info record-info">
        <span>
            <a href="http://beian.miit.gov.cn/" target="_blank">
                {{ profileStore.websiteRecordInfo }}
            </a>
        </span>
    </div>
    <div class="info system-info">
        <span>Powered by 
            <a href="https://github.com/m4ru1/Mollia" target="_blank">Mollia</a>
        </span>
    </div>
    <div class="info copyright-info">
        <span>Copyright©{{ profileStore.sinceYear }}  
            <a :href="profileStore.githubAccount" target="_blank"> 
                {{ `  ${profileStore.userName}` }} 
            </a>
        </span>
    </div>
</div>
</template>

<script setup>
import { useProfileStore } from '../../stores/ProfileStore';

let profileStore = useProfileStore();

</script>

<style scoped>
.page-footer{
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    .info{
        text-align: center;
        margin-top: 2px;
    }
    .record-info, .copyright-info{
        & > span > a{
            color: var(--theme-text-color);
            text-decoration: none;
            &:hover{
                text-decoration: underline solid;
            }
        }
    }
    .system-info{
        & > span > a{
            color: var(--theme-primary-color);
            text-decoration: none;
            &:hover{
                text-decoration: underline solid var(--theme-primary-color);
            }
        }
    }
}
</style>