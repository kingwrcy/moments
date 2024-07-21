export type ResultVO<T> = {
    code: number
    message?: string
    data: T
}

export type LoginResp = {
    id: number
    token: string
    username: string
}

export type CommentVO = {
    id: number
    content: string
    username: string
    replyTo: string
    createdAt: string
    updatedAt: string
    memoId: number
    author: number
}


export type  MemoVO = {
    id: number
    content: string
    location: string
    imgs: string
    favCount: number
    userId: number
    createdAt: string
    updatedAt: string
    externalFavicon: string
    pinned: string
    ext: string
    externalTitle: string
    externalUrl: string
    showType: number
    user: UserVO,
    comments: Array<CommentVO>
}

export type UserVO = {
    id: number
    username: string
    nickname: string
    avatarUrl: string
    slogan: string
    coverUrl: string
}
export type SysConfigVO = {
    adminUserName: string,
    title: string,
    favicon: string,
    beiAnNo: string,
    css: string,
    js: string,
    enableS3: boolean
}

