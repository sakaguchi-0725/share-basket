import { Module } from '@nestjs/common'
import { ClientsModule, Transport } from '@nestjs/microservices'
import { join } from 'path'
import { PersonalShoppingService } from './personal-shopping.service'
import { PersonalShoppingResolver } from './personal-shopping.resovler'
import { AuthModule } from '../auth/auth.module'

@Module({
  imports: [
    ClientsModule.register([
      {
        name: 'PERSONAL_PACKAGE',
        transport: Transport.GRPC,
        options: {
          package: 'personal',
          protoPath: join(__dirname, '../../proto/personal/shopping.proto'),
          url: process.env.PERSONAL_SHOPPING_SERVICE_URL,
        },
      },
    ]),
    AuthModule,
  ],
  providers: [PersonalShoppingResolver, PersonalShoppingService],
})
export class PersonalShoppingModule {}
