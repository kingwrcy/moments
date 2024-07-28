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
    tags: string
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
    enableGoogleRecaptcha: boolean,
    googleSiteKey: string,
    enableComment: boolean,
    maxCommentLength: number,
    memoMaxHeight: number,
    commentOrder: 'desc' | 'asc',
    timeFormat: 'timeAgo' | 'time',
}


export type MetingJSDTO = {
    id: string | undefined
    api: string | undefined
    server: "netease" | "tencent" | "kugou" | "xiami" | "baidu" | undefined,
    type: "song" | "playlist" | "album" | "search" | "artist" | undefined
}

export type MetingMusicServer = Exclude<MetingJSDTO['server'], undefined>
export type MetingMusicType = Exclude<MetingJSDTO['type'], undefined>


export type ExtDTO = {
    music: MusicDTO,
    doubanBook: DoubanBook,
    doubanMovie: DoubanMovie,
    video: Video,
}

export type MusicDTO = {
    id?: string,
    server?: MetingMusicServer,
    type?: MetingMusicType,
    api?: string
}

export type DoubanBook = {
    url?: string
    id?: string
    title?: string
    desc?: string
    image?: string
    isbn?: string
    author?: string
    rating?: string
    pubDate?: string
    keywords?: string
}

export type DoubanMovie = {
    url?: string
    id?: string
    title?: string
    desc?: string
    image?: string
    director?: string
    releaseDate?: string
    rating?: string
    actors?: string
    runtime?: string
}

export type Video = {
    type: 'youtube' | 'bilibili' | 'online'
    value: string
}

export type VideoType = Video["type"]
