using System;
using System.IO;
using ResourceRegistrator.Manager;
using ResourceRegistrator.Model;

class Program
{
    static void MenuPrint()
    {
        Console.WriteLine("Меню:");
        Console.WriteLine("\t1) Создать ресурсы");
        Console.WriteLine("\t2) Загрузить значения ресурсов из файла");
        Console.WriteLine("\t3) Добавить запрос");
        Console.WriteLine("\t4) Освободить ресурс");
        Console.WriteLine("\t5) Такт");
        Console.WriteLine("\t6) Информация о ресурсах");
        Console.WriteLine("\t7) Информация о запросах");
        Console.WriteLine("\t0) Выход");
    }

    static void MenuHandle(Manager manager, int option)
    {
        switch (option)
        {
            case 1:
                Console.Write("\tВведите число ресурсов: ");
                if (int.TryParse(Console.ReadLine(), out int resCnt))
                {
                    manager.InitModel(resCnt);
                }
                else
                {
                    Console.WriteLine("Некорректный ввод, не число.");
                }
                break;
            case 2:
                Console.Write("\tВведите путь до файла: ");
                var filePath = Console.ReadLine();
                try
                {
                    manager.InitFromFile(filePath);
                }
                catch (Exception ex)
                {
                    Console.WriteLine(ex.Message);
                }
                break;
            case 3:
                if (manager.CheckModel())
                {
                    Console.WriteLine("Ресурсы не созданы");
                    return;
                }
                Console.Write("\tВведите название ресурса: ");
                var resId = Console.ReadLine();
                Console.Write("\tВведите продолжительность запроса: ");
                if (int.TryParse(Console.ReadLine(), out int reqT))
                {
                    if (reqT > 0)
                    {
                        try
                        {
                            manager.AddRequest(new Request(resId, reqT));
                        }
                        catch (Exception ex)
                        {
                            Console.WriteLine(ex.Message);
                        }
                    }
                    else
                    {
                        Console.WriteLine("Некорректное время.");
                    }
                }
                else
                {
                    Console.WriteLine("Некорректный ввод, не число.");
                }
                break;
            case 4:
                if (manager.CheckModel())
                {
                    Console.WriteLine("Ресурсы не созданы");
                    return;
                }
                Console.Write("\tВведите название ресурса: ");
                resId = Console.ReadLine();
                try
                {
                    manager.FreeResource(resId);
                }
                catch (Exception ex)
                {
                    Console.WriteLine(ex.Message);
                }
                break;
            case 5:
                if (manager.CheckModel())
                {
                    Console.WriteLine("Ресурсы не созданы");
                    return;
                }
                manager.Work();
                break;
            case 6:
                if (manager.CheckModel())
                {
                    Console.WriteLine("Ресурсы не созданы");
                    return;
                }
                Console.WriteLine(manager.ModelInfo());
                break;
            case 7:
                if (manager.CheckModel())
                {
                    Console.WriteLine("Ресурсы не созданы");
                    return;
                }
                Console.WriteLine(manager.RequestsInfo());
                break;
            case 0:
                Console.WriteLine("Завершение работы...");
                Environment.Exit(0);
                break;
            default:
                Console.WriteLine("Некорректный ввод, неизвестная опция.");
                break;
        }
    }

    static void Main()
    {
        var manager = new Manager(new Model());
        while (true)
        {
            MenuPrint();
            Console.Write("Выберите опцию: ");
            if (int.TryParse(Console.ReadLine(), out int option))
            {
                MenuHandle(manager, option);
            }
            else
            {
                Console.WriteLine("Некорректный ввод, не число.");
            }
        }
    }
}
