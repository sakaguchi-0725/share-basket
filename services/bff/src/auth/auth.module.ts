import { Module } from '@nestjs/common'
import { AuthResolver } from './auth.resolver'
import { AuthService } from './auth.service'
import { ClientsModule, Transport } from '@nestjs/microservices'
import { join } from 'path'
import { AuthGuard } from './auth.guard'

@Module({
  imports: [
    ClientsModule.register([
      {
        name: 'AUTH_PACKAGE',
        transport: Transport.GRPC,
        options: {
          package: 'personal',
          protoPath: join(__dirname, '../../proto/personal/auth.proto'),
          url: process.env.PERSONAL_SHOPPING_SERVICE_URL,
        },
      },
    ]),
  ],
  providers: [AuthResolver, AuthService, AuthGuard],
  exports: [AuthGuard, AuthService],
})
export class AuthModule {}
