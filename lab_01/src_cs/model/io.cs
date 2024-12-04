using System;
using System.Collections.Generic;
using System.IO;

namespace ResourceRegistrator.Model
{
    public static class Io
    {
        public static void InitFromFile(this Model model, string filePath)
        {
            using (var file = new StreamReader(filePath))
            {
                string line;
                while ((line = file.ReadLine()) != null)
                {
                    var parts = line.Split(" ");
                    if (parts.Length != 2)
                    {
                        throw new Exception(Errors.IncorrectFile);
                    }
                    var resourceTag = parts[0].Trim();
                    var resourceStatus = parts[1].Trim();

                    ResourceStatus status;
                    switch (resourceStatus)
                    {
                        case nameof(ResourceStatus.Free):
                            status = ResourceStatus.Free;
                            break;
                        case nameof(ResourceStatus.Occupied):
                            status = ResourceStatus.Occupied;
                            break;
                        default:
                            throw new Exception(Errors.IncorrectFile);
                    }
                    model.AddResource(resourceTag, status);
                }
            }
        }
    }
}
