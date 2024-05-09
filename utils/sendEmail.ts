import nodemailer from "nodemailer";
import type { PrivateConfig } from "~/lib/types";
import fs from "fs/promises";

type SendEmailOptions = {
  email: string;
  subject: string;
  message: string;
};

export async function sendEmail(options: SendEmailOptions) {
  const config = (await fs.readFile(`${process.env.CONFIG_FILE}`)).toString();
  const privateConfig = JSON.parse(config).private as PrivateConfig;
  const transporter = nodemailer.createTransport({
    host: privateConfig.emailHost,
    port: privateConfig.emailPort || 587,
    secure: privateConfig.emailSecure || false,
    auth: {
      user: privateConfig.emailLoginName,
      pass: privateConfig.emailPassword,
    },
  });

  const mailOptions = {
    from: `"${privateConfig.emailFromName}" <${privateConfig.emailFrom}>`,
    to: options.email,
    subject: options.subject,
    text: options.message,
    html: options.message,
  };

  try {
    const info = await transporter.sendMail(mailOptions);
    return { success: true, messageId: info.messageId };
  } catch (error: any) {
    return { success: false, error: error.message };
  }
}
