export type Memo = {
  id: number;
  content: string;
  imgs: string;
  favCount: number;
  commentCount: number;
  userId: number;
  createdAt: string;
  updatedAt: string;
  user: User;
  music163Url?: string;
  bilibiliUrl?: string;
  location?: string;
  externalUrl?: string;
  externalTitle?: string;
  externalFavicon?: string;
  comments: Array<Comment>;
  pinned: boolean;
  ext: string;
  _count: {
    comments: number;
  };
  showType: number;
};

export type Comment = {
  id: number;
  content: string;
  replyTo: string;
  username: any;
  email: any;
  website: any;
  createdAt: string;
  updatedAt: string;
  memoId: number;
  author?: Boolean;
};

export type User = {
  username: string;
  nickname: string;
  slogan: string;
  id: number;
  avatarUrl: string;
  coverUrl: string;
  favicon: string;
  title: string;
  css: string;
  js: string;
  beianNo: string;
  enableS3: boolean;
  thumbnailSuffix:string;
};

export type DoubanBook = {
  title: string;
  desc: string;
  image: string;
  author: string;
  isbn: string;
  rating: string;
  url: string;
  pubDate: string;
  id: number;
};

export type DoubanMovie = {
  title: string;
  desc: string;
  image: string;
  director: string;
  rating: string;
  url: string;
  releaseDate: string;
  actors: string;
  runtime: string;
  id: number;
};

export type MemoExt = {
  doubanBook: DoubanBook;
  doubanMovie: DoubanMovie;
  youtubeUrl: string;
  videoUrl: string;
  localVideoUrl: string;
  audioUrl: string;
};

export type SysConfig = {
  private: {
    commentOrderBy: "desc";
    enableDouban: true;
    enableMusic163: true;
    enableVideo: true;
    googleRecaptchaSecretKey: "";
    googleRecaptchaEnable: false;
    enableNotifyByEmail: false;
    adminEmail: "";
    emailHost: "";
    emailPort: 587;
    emailSecure: true;
    emailLoginName: "";
    emailPassword: "";
    emailFrom: "";
    emailFromName: "";
    enableAliyunJudge: false;
    aliyunAk: "";
    aliyunSk: "";
  };
  public: {
    siteUrl: string;
    enableComment: boolean;
    enableShowComment: boolean;
    commentMaxLength: number;
    memoMaxLine: number;
    googleRecaptchaSiteKey: string;
    pageSize: number;
    dateTimeFormat: string;
    tencentMapKey: string;
  };
};

export type PublicConfig = SysConfig['public'];
export type PrivateConfig = SysConfig['private'];