
export default defineEventHandler(async (event) => {
  return {
    version: process.env.VERSION,
  };
});
