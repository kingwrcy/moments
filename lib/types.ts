export type Memo = {
  id: number;
  content: string;
  imgs: any;
  favCount: number;
  commentCount: number;
  userId: number;
  createdAt: string;
  updatedAt: string;
  user: User;
  comments: Array<Comment>;
  _count: {
    comments: number;
  };
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
}

export type User = {
  username: string;
  nickname: string;
  slogan: string;
  id: number;
  avatarUrl: string;
  coverUrl: string;
}