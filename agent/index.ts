// import amqp from 'amqplib/callback_api'
import amqp from 'amqplib'
import { applicationConfig } from 'configuration.js';
import { exit } from 'process';

const createChannelAsync = async (url: string): Promise<amqp.Channel | null> => {
    try {
        const connection: amqp.ChannelModel = await amqp.connect(url);
        const channel: amqp.Channel = await connection.createChannel();
        return channel;
    } catch (error) {
        console.error(error);
        return null;
    }
};

const assertQueueAsync = async (
    channel: amqp.Channel,
    queueName: string
): Promise<boolean> => {
    try {
        await channel.assertQueue(queueName, { durable: true });
        return true;
    } catch (error) {
        console.error(error);
        return false;
    }
};

const healthStatusUpdate = (channel: amqp.Channel, queue: string) => {
    channel.sendToQueue(
        queue,
        Buffer.from(JSON.stringify({ message: "Service is running normally" })),
        {
            contentType: "application/json",
            mandatory: true,
        }
    );
};

const onMessage = async (msg: amqp.ConsumeMessage | null): Promise<void> => {
    if (msg)
        console.log(" [x] Received %s", msg.content.toString());
}

(async () => {
    let timer: NodeJS.Timeout;
    const channel = await createChannelAsync(applicationConfig.mqQueueUrl);
    if (!channel) {
        exit(1);
    }
    await assertQueueAsync(channel, applicationConfig.mqHertbeatUrl);


    timer = setInterval(() => healthStatusUpdate(channel, applicationConfig.mqHertbeatUrl), 10 * 1000);

    channel.consume(applicationConfig.mqRequestUrl, onMessage, { noAck: true })
})();