import dotenv from 'dotenv';
import envvar from 'env-var';
dotenv.config();

const applicationConfig = {
  browserPath: envvar.get('PFLOW_BROWSER_PATH').required().asString(),
  mqQueueUrl: envvar.get('PFLOW_MQ_URL').required().asUrlString(),
  mqRequestUrl: envvar.get('PFLOW_MQ_REQUEST_URL').required().asUrlString(),
  mqHertbeatUrl: envvar.get('PFLOW_MQ_HEARTBEAT_URL').default('heartbeat').asUrlString()
};

export { applicationConfig };
