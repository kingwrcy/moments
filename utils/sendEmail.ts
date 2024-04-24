import nodemailer from 'nodemailer';

type SendEmailOptions = {
    email: string;
    subject: string;
    message: string;
};

export async function sendEmail(options: SendEmailOptions) {
    const transporter = nodemailer.createTransport({
        host: process.env.MAIL_HOST,
        port: parseInt(process.env.MAIL_PORT || '587'),
        secure: process.env.MAIL_SECURE === "true",
        auth: {
            user: process.env.MAIL_FROM,
            pass: process.env.MAIL_PASSWORD,
        },
    });

    const mailOptions = {
        from: `"${process.env.MAIL_NAME}" <${process.env.MAIL_FROM}>`,
        to: options.email,
        subject: options.subject,
        text: options.message,
        html: options.message,
    };

    try {
        const info = await transporter.sendMail(mailOptions);
        return { success: true, messageId: info.messageId };
    } catch (error) {
        return { success: false, error: error.message };
    }
}