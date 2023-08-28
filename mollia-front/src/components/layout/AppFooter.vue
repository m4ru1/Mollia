<template>
    <div class="app-footer">
        <a class="footer-content security-record">
            鲁ICP备2022040119号
        </a>
        <a class="footer-content tech-supported">
            Powered by <span style="color:#10b981;">Mollia</span>
        </a>
        <a class="footer-content copyright">
            Copyright from <span style="color:#10b981;">mollia</span> Since 2023
        </a>
    </div>
</template>

<style>
.app-footer{
    display: flex;
    flex-direction: column;
    justify-content: center;
    background-color: #111;
}

.footer-content{
    color: white;
    margin: 5px auto 5px auto;
    user-select: none;
}

.security-record{
    font-size: 15px;
}

</style>