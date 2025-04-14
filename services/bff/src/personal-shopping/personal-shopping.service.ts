import { Inject, Injectable, OnModuleInit } from '@nestjs/common'
import { ClientGrpc } from '@nestjs/microservices'
import {
  PersonalShoppingServiceClient,
  ShoppingStatus,
} from 'src/gen/personal/shopping'
import { ShoppingItem, Status } from './dto/shopping-item.dto'
import { CreateShoppingItemInput } from './dto/create.input'
import { firstValueFrom } from 'rxjs'
import { safeParseNumber } from 'src/common/utils'

@Injectable()
export class PersonalShoppingService implements OnModuleInit {
  private service: PersonalShoppingServiceClient

  constructor(@Inject('PERSONAL_PACKAGE') private client: ClientGrpc) {}

  onModuleInit() {
    this.service = this.client.getService<PersonalShoppingServiceClient>(
      'PersonalShoppingService',
    )
  }

  async create(input: CreateShoppingItemInput): Promise<ShoppingItem> {
    const observable = this.service.create({
      name: input.name,
      categoryId: String(input.categoryId),
    })

    const response = await firstValueFrom(observable)

    return {
      id: safeParseNumber(response.id),
      name: response.name,
      status: this.convertGqlStatus(response.status),
      categoryId: safeParseNumber(response.categoryId),
    }
  }

  async getAll(): Promise<ShoppingItem[]> {
    const observable = this.service.getAll({})

    const response = await firstValueFrom(observable)

    return response.items.map((item) => {
      return {
        id: safeParseNumber(item.id),
        name: item.name,
        status: this.convertGqlStatus(item.status),
        categoryId: safeParseNumber(item.categoryId),
      }
    })
  }

  private convertGqlStatus(status: ShoppingStatus): Status {
    switch (status) {
      case ShoppingStatus.STATUS_UNPURCHASED:
        return Status.UNPURCHASED
      case ShoppingStatus.STATUS_PURCHASED:
        return Status.PURCHASED
      default:
        return Status.UNPURCHASED
    }
  }
}
