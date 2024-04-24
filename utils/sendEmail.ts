import nodemailer from 'nodemailer';

type SendEmailOptions = {
    email: string;
    subject: string;
    message: string;
};

export async function sendEmail(options: SendEmailOptions) {
    const config = useRuntimeConfig()
    const transporter = nodemailer.createTransport({
        host: config.mailHost,
        port: config.mailPort || 587,
        secure: config.mailSecure || false,
        auth: {
            user: config.mailFrom,
            pass: config.mailPassword,
        },
    });

    const mailOptions = {
        from: `"${config.mailName}" <${config.mailFrom}>`,
        to: options.email,
        subject: options.subject,
        text: options.message,
        html: options.message,
    };

    try {
        const info = await transporter.sendMail(mailOptions);
        return { success: true, messageId: info.messageId };
    } catch (error:any) {
        return { success: false, error: error.message };
    }
}