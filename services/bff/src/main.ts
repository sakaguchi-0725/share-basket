import { NestFactory } from '@nestjs/core'
import { AppModule } from './app.module'
import { GrpcExceptionFilter } from './common/error'
import * as cookieParser from 'cookie-parser'

async function bootstrap() {
  const app = await NestFactory.create(AppModule)
  app.useGlobalFilters(new GrpcExceptionFilter())
  app.use(cookieParser())

  app.enableCors({
    origin: true, // 開発環境では全オリジンを許可
    credentials: true,
  })

  await app.listen(process.env.PORT ?? 3000)
}
void bootstrap()
